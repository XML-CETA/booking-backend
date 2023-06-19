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

func (repo *ReservationMongoDBStore) CreateReservation(reservation domain.Reservation) error {
	_, err := repo.reservations.InsertOne(context.Background(), reservation)

	return err
}

func (repo *ReservationMongoDBStore) GetFirstActive(accommodation string, dateFrom, dateTo string) (domain.Reservation, error) {
	filter := bson.D{
		{Key: "accommodation", Value: accommodation},
		{Key: "datefrom", Value: dateFrom},
		{Key: "dateto", Value: dateTo},
		{Key: "status", Value: domain.Reserved},
	}

	var result domain.Reservation
	err := repo.reservations.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (repo *ReservationMongoDBStore) GetById(reservation primitive.ObjectID) (domain.Reservation, error) {
	filter := bson.D{
		{Key: "_id", Value: reservation},
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

func (repo *ReservationMongoDBStore) GetWaitingReservations(host string) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	filter := bson.D{
		{Key: "host", Value: host},
		{Key: "status", Value: domain.Waiting},
	}
	result, err := repo.reservations.Find(context.Background(), filter)
	for result.Next(context.TODO()) {
		var reservation domain.Reservation
		err := result.Decode(&reservation)
		if err == nil {
			reservations = append(reservations, reservation)
		}
	}
	return reservations, err
}

func (repo *ReservationMongoDBStore) GetByUserAndHost(user, host string) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	filter := bson.D{
		{Key: "user", Value: user},
		{Key: "host", Value: host},
		{Key: "status", Value: domain.Reserved},
	}
	result, err := repo.reservations.Find(context.Background(), filter)
	for result.Next(context.TODO()) {
		var reservation domain.Reservation
		err := result.Decode(&reservation)
		if err == nil {
			reservations = append(reservations, reservation)
		}
	}
	return reservations, err
}

func (repo *ReservationMongoDBStore) GetByIdAndUser(reservation primitive.ObjectID, user string) (domain.Reservation, error) {
	filter := bson.D{
		{Key: "_id", Value: reservation},
		{Key: "user", Value: user},
	}

	var result domain.Reservation
	err := repo.reservations.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (repo *ReservationMongoDBStore) GetWaitingByAccommodation(accommodation string) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	filter := bson.D{
		{Key: "accommodation", Value: accommodation},
		{Key: "status", Value: domain.Waiting},
	}
	result, err := repo.reservations.Find(context.Background(), filter)
	for result.Next(context.TODO()) {
		var reservation domain.Reservation
		err := result.Decode(&reservation)
		if err == nil {
			reservations = append(reservations, reservation)
		}
	}
	return reservations, err
}

func (repo *ReservationMongoDBStore) ConfirmReservation(reservation primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: reservation}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: domain.Reserved},
		}},
	}
	_, err := repo.reservations.UpdateOne(context.Background(), filter, update)
	return err
}

func (repo *ReservationMongoDBStore) Cancel(reservation primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: reservation}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: domain.Canceled},
		}},
	}
	_, err := repo.reservations.UpdateOne(context.Background(), filter, update)
	return err
}

func (repo *ReservationMongoDBStore) CountCanceled(host string) (int32, error) {
	filterStatus := bson.D{{Key: "status", Value: domain.Canceled}}
	filter := makeStatusHostFilter(filterStatus, host)

	return repo.countDocuments(filter)
}

func (repo *ReservationMongoDBStore) CountUserCanceled(user string) (int32, error) {
	filterStatus := bson.D{{Key: "status", Value: domain.Canceled}}
	filter := makeStatusUserFilter(filterStatus, user)
	return repo.countDocuments(filter)
}

func (repo *ReservationMongoDBStore) CountExpired(host string) (int32, error) {
	filterStatus := bson.D{{Key: "status", Value: domain.Expired}}
	filter := makeStatusHostFilter(filterStatus, host)

	return repo.countDocuments(filter)
}


func (repo *ReservationMongoDBStore) CountActive(user string, role string) (int32, error) {
	filterStatus := bson.D{{Key: "status", Value: domain.Reserved}}
	var filter primitive.D
  if role == "HOST" {
    filter = makeStatusHostFilter(filterStatus, user)
  } else {
    filter = makeStatusUserFilter(filterStatus, user)
  }

	return repo.countDocuments(filter)
}

func (repo *ReservationMongoDBStore) CountNonCanceled(host string) (int32, error) {
	filterStatus := bson.D{{
		Key: "status", Value: bson.D{{Key: "$ne", Value: domain.Canceled}},
	}}

	filter := makeStatusHostFilter(filterStatus, host)

	return repo.countDocuments(filter)
}

func (repo *ReservationMongoDBStore) countDocuments(filter primitive.D) (int32, error) {
	count, err := repo.reservations.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}

func (repo *ReservationMongoDBStore) GetHostIntervalSum(host string) (int32, error) {
	matchStage := bson.D{{
		Key: "$match", Value: bson.D{{Key: "host", Value: host}},
	}}

	groupStage := bson.D{{
		Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$host"},
			{Key: "sum", Value: bson.D{{Key: "$sum", Value: "$duration"}}},
		},
	}}

	cursor, err := repo.reservations.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results[0]["sum"].(int32), nil
}

func makeStatusHostFilter(filterStatus primitive.D, host string) primitive.D {
	return bson.D{{
		Key: "$and", Value: bson.A{
			filterStatus,
			bson.D{{Key: "host", Value: host}},
		},
	}}
}

func makeStatusUserFilter(filterStatus primitive.D, host string) primitive.D {
	return bson.D{{
		Key: "$and", Value: bson.A{
			filterStatus,
			bson.D{{Key: "user", Value: host}},
		},
	}}
}
