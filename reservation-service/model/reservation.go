package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status int32

const (
	Waiting Status = iota
	Reserved
	Expired
)

type Reservation struct{
	id primitive.ObjectID `bson:"_id,omitempty"`
	accommodation int32
	offer int32
	dateFrom string
	dateTo string
	guests int32
	Status Status
}


func MakeReservation(accommodation, offer, guests int32, dateFrom, dateTo string) Reservation {
	return Reservation{
		accommodation:  accommodation,
		offer: offer,
		guests: guests,
		dateFrom: dateFrom,
		dateTo: dateTo,
	}
}

