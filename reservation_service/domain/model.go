package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status int32

const (
	Waiting Status = iota
	Reserved
	Expired
  Canceled
)

type Reservation struct{
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Accommodation string
	User string
	DateFrom string
	DateTo string
	Guests int32
	Status Status
  Host string
}

func MakeReservation(guests int32, accommodation, user, dateFrom, dateTo string) Reservation {
	return Reservation{
		Accommodation:  accommodation,
		User: user,
		Guests: guests,
		DateFrom: dateFrom,
		DateTo: dateTo,
	}
}
