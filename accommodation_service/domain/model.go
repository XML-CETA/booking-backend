package domain

import (
	pb "booking-backend/common/proto/accommodation_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id        primitive.ObjectID   `json:"id,omitempty"  bson:"_id,omitempty"`
	Longitude float64              `json:"longitude" bson:"longitude"`
	Latitude  float64              `json:"latitude" bson:"latitude"`
	Address   AccommodationAddress `json:"address" bson"address"`
	MinGuests int32                `json:"minGuests" bson:"minGuests"`
	MaxGuests int32                `json:"maxGuests" bson:"maxGuests"`
	Name      string               `json:"name" bson:"name"`
}

type AccommodationAddress struct {
	Street  string `json:"street" bson:"street"`
	Number  int32  `json:"number" bson:"number"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
}

func MakeCreateAccommodation(accommodation *pb.AccommodationCreateRequest) Accommodation {
	address := AccommodationAddress{
		Street:  accommodation.Address.Street,
		Number:  accommodation.Address.Number,
		City:    accommodation.Address.City,
		Country: accommodation.Address.Country,
	}

	return Accommodation{
		Longitude: accommodation.Longitude,
		Latitude:  accommodation.Latitude,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Name:      accommodation.Name,
		Address:   address,
	}
}

func MakeAccommodation(accommodation *pb.SingleAccommodation) Accommodation {

	address := AccommodationAddress{
		Street:  accommodation.Address.Street,
		Number:  accommodation.Address.Number,
		City:    accommodation.Address.City,
		Country: accommodation.Address.Country,
	}

	id, _ := primitive.ObjectIDFromHex(accommodation.Id)

	return Accommodation{
		Id:        id,
		Longitude: accommodation.Longitude,
		Latitude:  accommodation.Latitude,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Name:      accommodation.Name,
		Address:   address,
	}
}
