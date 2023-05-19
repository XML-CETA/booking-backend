package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	GetAll() ([]Reservation, error)
	CreateReservation(reservation Reservation) error
	GetFirstActive(accommodation string, dateFrom, dateTo string) (Reservation, error)
	GetById(reservation primitive.ObjectID) (Reservation, error)
	Delete(reservation primitive.ObjectID) (error)
}
