package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status int32

const (
	Waiting Status = iota
	Reserved
	Expired
)

type Reservation struct{
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Accommodation int32
	Offer int32
	DateFrom string
	DateTo string
	Guests int32
	Status Status
}

func MakeReservation(accommodation, offer, guests int32, dateFrom, dateTo string) Reservation {
	return Reservation{
		Accommodation:  accommodation,
		Offer: offer,
		Guests: guests,
		DateFrom: dateFrom,
		DateTo: dateTo,
	}
}
