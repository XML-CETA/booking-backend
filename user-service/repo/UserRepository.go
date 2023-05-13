package repo

import (
	"context"
	"example/grpc/model"
	"example/grpc/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
}

func (repo *UserRepository) Create(flight *model.User) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("users").Collection("users")
	_, err := coll.InsertOne(context.TODO(), flight)

	return err
}

func (repo *UserRepository) Delete(email string) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("users").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var result model.User
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		_, err2 := coll.DeleteOne(context.TODO(), result)
		return err2
	}

	return err
}

func (repo *UserRepository) Update(user *model.User) error {
	client, cancel := utils.GetConn()
	defer cancel()

	filter := bson.D{{Key: "email", Value: user.Email}}

	coll := client.Database("users").Collection("users")
	_, err := coll.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: user}})

	return err
}

func (repo *UserRepository) GetOne(email string) (model.User, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("users").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var result model.User
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
