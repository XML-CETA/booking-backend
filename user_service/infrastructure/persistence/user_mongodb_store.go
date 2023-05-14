package persistence

import (
	"context"

	"booking-backend/user-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "users"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	userss := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: userss,
	}
}

func (store *UserMongoDBStore) Create(user domain.User) error {
	_, err := store.users.InsertOne(context.TODO(), user)

	return err
}

func (store *UserMongoDBStore) Delete(email string) error {
	filter := bson.D{{Key: "email", Value: email}}

	var result *domain.User
	err := store.users.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		_, err2 := store.users.DeleteOne(context.TODO(), result)
		return err2
	}

	return err
}

func (store *UserMongoDBStore) Update(user domain.User) error {

	filter := bson.D{{Key: "email", Value: user.Email}}

	_, err := store.users.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: user}})

	return err
}

func (store *UserMongoDBStore) GetOne(email string) (domain.User, error) {

	filter := bson.D{{Key: "email", Value: email}}

	var result domain.User
	err := store.users.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
