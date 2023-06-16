package persistence

import (
	"booking-backend/rating-service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (store *RatingAccommodationMongoDBStore) Create(rate domain.RatingAccommodation) error {
	_, err := store.ratings.InsertOne(context.TODO(), rate)
	return err
}

func (store *RatingAccommodationMongoDBStore) Update(rate domain.RatingAccommodation) error {
	filter := bson.D{{Key: "accommodation", Value: rate.Accommodation}, {Key: "user", Value: rate.User}}

	_, err := store.ratings.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: rate}})

	return err
}

func (store *RatingAccommodationMongoDBStore) Delete(rate domain.RatingAccommodation) error {
	filter := bson.D{{Key: "accommodation", Value: rate.Accommodation}, {Key: "user", Value: rate.User}}

	_, err := store.ratings.DeleteOne(context.TODO(), filter)
	return err
}

func (store *RatingAccommodationMongoDBStore) GetByUserAndAccommodation(userEmail, accommodation string) (domain.RatingAccommodation, error) {
	filter := bson.D{{Key: "accommodation", Value: accommodation}, {Key: "user", Value: userEmail}}

	var result domain.RatingAccommodation
	err := store.ratings.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

func (store *RatingAccommodationMongoDBStore) GetAllByAccommodation(accommodation string) ([]domain.RatingAccommodation, error) {
	var rates []domain.RatingAccommodation
	filter := bson.D{{Key: "accommodation", Value: accommodation}}

	dataResult, err := store.ratings.Find(context.TODO(), filter)
	for dataResult.Next(context.TODO()) {
		var rate domain.RatingAccommodation
		err := dataResult.Decode(&rate)
		if err == nil {
			rates = append(rates, rate)
		}
	}

	return rates, err
}
