package service

import (
	"errors"
	"example/grpc/model"
	"example/grpc/proto/reservation"
	"example/grpc/repo"
)

type Service struct {
	Repo *repo.Repository
}

func (service *Service) CreateReservation(reservation model.Reservation) error {
	reservation.Status = model.Reserved

	_, err := service.Repo.GetFirstByDates(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if (err == nil) {
		return errors.New("Could not create, reservation with the same interval already exists")
	}

	// TODO: ask accommodation for status

	return service.Repo.CreateReservation(reservation)
}

func (service *Service) GetAll() ([]model.Reservation, error) {
	return service.Repo.GetAll()
}

func (service *Service) ConvertToGrpcList(reservations []model.Reservation) []*reservation.Reservation {
	var converted []*reservation.Reservation

	for _, entity := range reservations {
		newRes := reservation.Reservation{
			Id: entity.Id.Hex(),
			Accommodation: entity.Accommodation,
			DateFrom: entity.DateFrom,
			DateTo: entity.DateTo,
			Guests: entity.Guests,
			Offer: entity.Offer,
			Status: int32(entity.Status),
		}

		converted = append(converted, &newRes)
	}

	return converted
}
