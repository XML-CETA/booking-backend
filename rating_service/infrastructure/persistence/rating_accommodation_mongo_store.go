package persistence

import (
	"booking-backend/rating-service/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "ratings"
	COLLECTION = "accommodations"
)

type RatingAccommodationMongoDBStore struct {
	ratings *mongo.Collection
}

func NewRatingAccommodationMongoDBStore(client *mongo.Client) domain.RatingAccommodationStore {
	ratings := client.Database(DATABASE).Collection(COLLECTION)
	return &RatingAccommodationMongoDBStore{
		ratings: ratings,
	}
}
