package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int32

const (
	Waiting Status = iota
	Reserved
	Expired
	Canceled
)

type Reservation struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Accommodation string
	User          string
	DateFrom      string
	DateTo        string
	Guests        int32
	Status        Status
	Host          string
	Duration      int32
}

type Flight struct {
	Id               primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	FlighDateAndTime time.Time          `json:"dateTime"`
	StartingPoint    string             `json:"startingPoint"`
	Destination      string             `json:"destination"`
	Price            int32              `json:"price"`
	Seats            int32              `json:"allSeats"`
	RemainingSeats   int32              `json:"remainingSeats,omitempty"`
}

func MakeReservation(guests int32, accommodation, user, dateFrom, dateTo string) Reservation {
	return Reservation{
		Accommodation: accommodation,
		User:          user,
		Guests:        guests,
		DateFrom:      dateFrom,
		DateTo:        dateTo,
	}
}
