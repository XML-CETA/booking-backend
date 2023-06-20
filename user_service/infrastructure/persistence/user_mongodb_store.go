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

		_, err := store.users.DeleteOne(context.TODO(), filter)
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

func (store *UserMongoDBStore) UpdateProminent(isProminent bool, host string) error {
	filter := bson.D{{Key: "email", Value: host}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "isprominent", Value: isProminent},
		}},
	}

	_, err := store.users.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *UserMongoDBStore) GetAllProminent() ([]string, error) {
	filter := bson.D{{Key: "isprominent", Value: true}}

	cur, err := store.users.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var emails []string
	for cur.Next(context.TODO()) {
		var user domain.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		emails = append(emails, user.Email)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}
