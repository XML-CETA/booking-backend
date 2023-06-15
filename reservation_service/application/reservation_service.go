package application

import (
	"booking-backend/common/clients"
	"booking-backend/common/messaging"
	"booking-backend/common/proto/accommodation_service"
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/domain"
	"booking-backend/reservation-service/startup/config"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store domain.ReservationStore
  prominentHostPublisher messaging.PublisherModel
}

func NewReservationService(store domain.ReservationStore, prominentHostPublisher messaging.PublisherModel) *ReservationService {
	return &ReservationService{
		store: store,
    prominentHostPublisher: prominentHostPublisher,
	}
}

func (service *ReservationService) CreateReservation(reservation domain.Reservation) error {
	_, err := service.store.GetFirstActive(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		return errors.New("Could not create, an active reservation with the same interval already exists")
	}

	// TODO: Ask accommodation for status, can use validate reservation method
	reservation.Status = domain.Reserved

	_, err = validateReservation(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err != nil {
		return err
	}

	return service.store.CreateReservation(reservation)
}

func (service *ReservationService) GetAll() ([]domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) Delete(reservationId string, user string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}

	reservation, err := service.store.GetByIdAndUser(id, user)

	if reservation.Status == domain.Waiting {
		return service.store.Delete(id)
	}

	if !canDeleteReservation(reservation) {
		return errors.New("Can not delete, reservation starts in less then a day or is in progress")
	}

	return service.store.Delete(id)
}

func (service *ReservationService) ConvertToGrpcList(reservations []domain.Reservation) []*pb.Reservation {
	var converted []*pb.Reservation

	for _, entity := range reservations {
		newRes := pb.Reservation{
			Id:            entity.Id.Hex(),
			Accommodation: entity.Accommodation,
			DateFrom:      entity.DateFrom,
			DateTo:        entity.DateTo,
			Guests:        entity.Guests,
			User:		   entity.User,
			Status:        int32(entity.Status),
		}

		converted = append(converted, &newRes)
	}

	return converted
}

func getAccommodationClient() accommodation_service.AccommodationServiceClient {
	return clients.NewAccommodationClient(fmt.Sprintf("%s:%s", config.NewConfig().AccommodationServiceHost, config.NewConfig().AccommodationServicePort))
}

func validateReservation(accommodationId, dateFrom, dateTo string) (*accommodation_service.ValidateReservationResponse, error) {
	accommodation := getAccommodationClient()

	return accommodation.ValidateReservation(context.Background(), &accommodation_service.ValidateReservationRequest{
		Accommodation: accommodationId,
		DateFrom: dateFrom,
		DateTo: dateTo,
	})
}

func checkReservationDate(dateFrom string) bool {
	date, err := time.Parse(time.DateOnly, dateFrom)
	if err != nil {
		return false
	}

	now := time.Now().UTC()
	return now.UTC().Compare(date.AddDate(0, 0, -1)) == -1
}

func canDeleteReservation(reservation domain.Reservation) bool {
	return reservation.Status == domain.Reserved && checkReservationDate(reservation.DateFrom)
}
