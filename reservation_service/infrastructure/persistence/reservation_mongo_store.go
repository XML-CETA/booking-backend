package persistence

import (
	"booking-backend/reservation-service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
const (
	DATABASE   = "scheduling"
	COLLECTION = "reservations"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (repo *ReservationMongoDBStore) CreateReservation(reservation domain.Reservation) (error) {
	_, err := repo.reservations.InsertOne(context.Background(), reservation)

	return err
}

func (repo *ReservationMongoDBStore) GetFirstActive(accommodation string, dateFrom, dateTo string) (domain.Reservation, error) {
	filter :=bson.D{
		{ Key: "accommodation", Value: accommodation },
		{ Key: "datefrom", Value: dateFrom },
		{ Key: "dateto", Value: dateTo },
		{ Key: "status", Value: domain.Reserved },
	}

	var result domain.Reservation
	err := repo.reservations.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (repo *ReservationMongoDBStore) GetAll() ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	dataResult, err := repo.reservations.Find(context.Background(), bson.M{})
	for dataResult.Next(context.TODO()) {
		var reservation domain.Reservation
		err := dataResult.Decode(&reservation)
		if err == nil {
			reservations = append(reservations, reservation)
		}
	}

	return reservations, err
}

func (repo *ReservationMongoDBStore) GetById(reservation primitive.ObjectID) (domain.Reservation, error) {
	filter :=bson.D{{ Key: "_id", Value: reservation }}

	var result domain.Reservation
	err := repo.reservations.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (repo *ReservationMongoDBStore) Delete(reservation primitive.ObjectID) error {
	filter :=bson.D{{ Key: "_id", Value: reservation }}

	_, err := repo.reservations.DeleteOne(context.Background(), filter)
	return err
}
