package application

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/accommodation_service"
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/domain"
	"booking-backend/reservation-service/startup/config"
	"context"
	"errors"
	"fmt"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) CreateReservation(reservation domain.Reservation) error {
	_, err := service.store.GetFirstActive(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		return errors.New("Could not create, an active reservation with the same interval already exists")
	}

	// TODO: Ask accommodation for status, can use validate reservation method
	reservation.Status = domain.Reserved

	accommodation := getAccommodationClient()

	_, err = accommodation.ValidateReservation(context.Background(), &accommodation_service.ValidateReservationRequest{
		Accommodation: reservation.Accommodation,
		DateFrom: reservation.DateFrom,
		DateTo: reservation.DateTo,
	})

	if err != nil {
		return err
	}

	return service.store.CreateReservation(reservation)
}

func (service *ReservationService) GetAll() ([]domain.Reservation, error) {
	return service.store.GetAll()
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
