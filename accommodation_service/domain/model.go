package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id        primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	Longitude float64            `json:"longitude" bson:"longitude"`
	Latitude  float64            `json:"latitude" bson:"latitude"`
	// Address   Address            `json:"address" bson"address"`
	MinGuests int32  `json:"minGuests" bson:"minGuests"`
	MaxGuests int32  `json:"maxGuests" bson:"maxGuests"`
	Name      string `json:"name" bson:"name"`
}

type Address struct {
	Street  string `json:"street" bson:"street"`
	Number  int32  `json:"number" bson:"number"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
}

func MakeAccommodation(longitude, latitude float64, minGuests, maxGuests int32, name string) Accommodation {
	return Accommodation{
		Longitude: longitude,
		Latitude:  latitude,
		MinGuests: minGuests,
		MaxGuests: maxGuests,
		Name:      name,
	}
}
