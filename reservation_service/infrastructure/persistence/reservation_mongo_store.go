package persistence

import (
	"booking-backend/reservation-service/domain"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
	result, err := repo.reservations.InsertOne(context.Background(), reservation)

	log.Println("Inserted with id: ", result.InsertedID)

	return err
}

func (repo *ReservationMongoDBStore) GetFirstByDates(accommodation int32, dateFrom, dateTo string) (domain.Reservation, error) {

	filter :=bson.D{
		{ Key: "accommodation", Value: accommodation },
		{ Key: "datefrom", Value: dateFrom },
		{ Key: "dateto", Value: dateTo },
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
