package service

import (
	"example/grpc/model"
	"example/grpc/repo"
)

type Service struct {
	Repo *repo.Repository
}

func (service *Service) CreateReservation(reservation model.Reservation) error {
	reservation.Status = model.Waiting
	return service.Repo.CreateReservation(reservation)
}
