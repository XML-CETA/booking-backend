package persistence

import (
	"booking-backend/user-service/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingMongoDBStore struct {
	ratings *mongo.Collection
}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingStore {
	ratingsCol := client.Database("users").Collection("ratings")
	return &RatingMongoDBStore{
		ratings: ratingsCol,
	}
}

func (store *RatingMongoDBStore) Create(rating domain.Rating) error {
	_, err := store.ratings.InsertOne(context.TODO(), rating)

	return err
}
