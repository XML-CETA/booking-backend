package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationStore interface {
	GetAll() ([]Reservation, error)
	CreateReservation(reservation Reservation) error
	GetFirstActive(accommodation string, dateFrom, dateTo string) (Reservation, error)
	GetByIdAndUser(reservation primitive.ObjectID, user string) (Reservation, error)
	Cancel(reservation primitive.ObjectID) (error)
  CountCanceled(host string) (int32, error)
  CountNonCanceled(host string) (int32, error)
  CountExpired(host string) (int32, error)
}
