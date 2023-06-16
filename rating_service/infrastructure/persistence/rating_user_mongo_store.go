package persistence

import (
	"booking-backend/rating-service/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingUserMongoDBStore struct {
	ratings *mongo.Collection
}

func NewRatingUserMongoDBStore(client *mongo.Client) domain.RatingUserStore {
	ratingsCol := client.Database("users").Collection("ratings")
	return &RatingUserMongoDBStore{
		ratings: ratingsCol,
	}
}

func (store *RatingUserMongoDBStore) Create(rating domain.RatingUser) error {
	_, err := store.ratings.InsertOne(context.TODO(), rating)

	return err
}
