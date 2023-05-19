package application

import (
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/domain"
	"errors"
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

	// TODO: ask accommodation for status
	reservation.Status = domain.Reserved

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
