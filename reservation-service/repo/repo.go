package repo

import (
	"context"
	"example/grpc/model"
	"example/grpc/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
}

func (repo *Repository) CreateReservation(reservation model.Reservation) (error) {
	client, cancel := utils.GetDbClient();
	defer cancel()

	coll := client.Database("reservations").Collection("reservations")
	result, err := coll.InsertOne(context.Background(), reservation)

	log.Println("Inserted with id: ", result.InsertedID)

	return err
}

func (repo *Repository) GetFirstByDates(accommodation int32, dateFrom, dateTo string) (model.Reservation, error) {
	client, cancel := utils.GetDbClient();
	defer cancel()

	coll := client.Database("reservations").Collection("reservations")

	filter :=bson.D{
		{ Key: "accommodation", Value: accommodation },
		{ Key: "datefrom", Value: dateFrom },
		{ Key: "dateto", Value: dateTo },
	}

	var result model.Reservation
	err := coll.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (repo *Repository) GetAll() ([]model.Reservation, error) {
	client, cancel := utils.GetDbClient()
	defer cancel()

	coll := client.Database("reservations").Collection("reservations")

	var reservations []model.Reservation
	dataResult, err := coll.Find(context.Background(), bson.M{})
	for dataResult.Next(context.TODO()) {
		var reservation model.Reservation
		err := dataResult.Decode(&reservation)
		if err == nil {
			reservations = append(reservations, reservation)
		}
	}

	return reservations, err
}
