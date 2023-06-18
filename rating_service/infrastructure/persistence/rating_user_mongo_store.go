package persistence

import (
	"booking-backend/rating-service/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingUserMongoDBStore struct {
	ratings *mongo.Collection
}

func NewRatingUserMongoDBStore(client *mongo.Client) domain.RatingUserStore {
	ratingsCol := client.Database("ratings").Collection("users")
	return &RatingUserMongoDBStore{
		ratings: ratingsCol,
	}
}

func (store *RatingUserMongoDBStore) Create(rating *domain.RatingUser) (primitive.ObjectID, error) {
	result, err := store.ratings.InsertOne(context.TODO(), rating)
	return result.InsertedID.(primitive.ObjectID), err
}

func (store *RatingUserMongoDBStore) GetHostRates(host string) ([]domain.RatingUser, error) {
	var ratings []domain.RatingUser
	filter := bson.D{
		{Key: "host", Value: host},
	}
	result, err := store.ratings.Find(context.Background(), filter)
	for result.Next(context.TODO()) {
		var rating domain.RatingUser
		err := result.Decode(&rating)
		if err == nil {
			ratings = append(ratings, rating)
		}
	}
	return ratings, err
}

func (store *RatingUserMongoDBStore) GetByHostAndUser(host, user string, id primitive.ObjectID) (domain.RatingUser, error) {
	filter := bson.D{
		{Key: "_id", Value: bson.D{{Key: "$ne", Value: id}}},
		{Key: "ratedUser", Value: host},
		{Key: "ratedBy", Value: user},
	}

	var result domain.RatingUser
	err := store.ratings.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (store *RatingUserMongoDBStore) UpdateStatus(id primitive.ObjectID, status domain.Status) error {
	result, err := store.ratings.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"status", status}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}
