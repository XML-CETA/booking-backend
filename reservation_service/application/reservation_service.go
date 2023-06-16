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
	store                  domain.ReservationStore
	prominentHostPublisher messaging.PublisherModel
}

func NewReservationService(store domain.ReservationStore, prominentHostPublisher messaging.PublisherModel) *ReservationService {
	return &ReservationService{
		store:                  store,
		prominentHostPublisher: prominentHostPublisher,
	}
}

func (service *ReservationService) CreateReservation(reservation domain.Reservation) error {
	_, err := service.store.GetFirstActive(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		return errors.New("Could not create, an active reservation with the same interval already exists")
	}

	response, err := validateReservation(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	reservation.Host = response.Host
	isAutomatic, err := isAutomaticConfirmation(reservation.Accommodation)
	if isAutomatic.IsAutomaticConfirmation {
		reservation.Status = domain.Reserved
	}
	if err != nil {
		return err
	}

	return service.store.CreateReservation(reservation)
}

func (service *ReservationService) GetAll() ([]domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) GetWaitingReservations(host string) ([]domain.Reservation, error) {
	return service.store.GetWaitingReservations(host)
}

func (service *ReservationService) Delete(reservationId string, user string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}

	reservation, err := service.store.GetByIdAndUser(id, user)

	if reservation.Status == domain.Waiting {
		return service.store.Cancel(id)
	}

	if !canDeleteReservation(reservation) {
		return errors.New("Can not delete, reservation starts in less then a day or is in progress")
	}

	return service.store.Cancel(id)
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
			User:          entity.User,
			Status:        int32(entity.Status),
			Host:          entity.Host,
		}

		converted = append(converted, &newRes)
	}

	return converted
}

func (service *ReservationService) ConvertToGrpcWaitingReservations(reservations []domain.Reservation) []*pb.WaitingReservation {
	var converted []*pb.WaitingReservation

	for _, entity := range reservations {
		canceledCount, _ := service.store.CountCanceled(entity.User)
		newRes := pb.WaitingReservation{
			Id:                             entity.Id.Hex(),
			Accommodation:                  entity.Accommodation,
			DateFrom:                       entity.DateFrom,
			DateTo:                         entity.DateTo,
			Guests:                         entity.Guests,
			User:                           entity.User,
			UserCanceledReservationsNumber: canceledCount,
		}
		converted = append(converted, &newRes)
	}
	return converted
}

func (service *ReservationService) getCancelRate(host string) (float32, error) {
	nonCanceled, err := service.store.CountNonCanceled(host)
	if err != nil {
		return 0.0, err
	}

	canceled, err := service.store.CountCanceled(host)
	if err != nil {
		return 0.0, err
	}

	return float32(nonCanceled / canceled), nil
}

func (service *ReservationService) getExpiredCount(host string) (int32, error) {
	return service.store.CountExpired(host)
}

func getAccommodationClient() accommodation_service.AccommodationServiceClient {
	return clients.NewAccommodationClient(fmt.Sprintf("%s:%s", config.NewConfig().AccommodationServiceHost, config.NewConfig().AccommodationServicePort))
}

func validateReservation(accommodationId, dateFrom, dateTo string) (*accommodation_service.ValidateReservationResponse, error) {
	accommodation := getAccommodationClient()

	return accommodation.ValidateReservation(context.Background(), &accommodation_service.ValidateReservationRequest{
		Accommodation: accommodationId,
		DateFrom:      dateFrom,
		DateTo:        dateTo,
	})
}

func isAutomaticConfirmation(accommodationId string) (*accommodation_service.IsAutomaticConfirmationResponse, error) {
	accommodation := getAccommodationClient()
	return accommodation.IsAutomaticConfirmation(context.Background(), &accommodation_service.AccommodationIdRequest{
		Id: accommodationId,
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
