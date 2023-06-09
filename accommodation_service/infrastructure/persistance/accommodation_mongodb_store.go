package persistance

import (
	"context"
	"log"

	"booking-backend/accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodations"
	COLLECTION = "accommodations"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) GetAll() ([]domain.Accommodation, error) {

	var accommodations []domain.Accommodation
	dataResult, err := store.accommodations.Find(context.Background(), bson.M{})
	for dataResult.Next(context.TODO()) {
		var accommodation domain.Accommodation
		err := dataResult.Decode(&accommodation)
		if err == nil {
			accommodations = append(accommodations, accommodation)
		}
	}

	return accommodations, err
}

func (store *AccommodationMongoDBStore) GetById(id primitive.ObjectID) (*domain.Accommodation, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var result domain.Accommodation
	err := store.accommodations.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}

func (store *AccommodationMongoDBStore) Create(accommodation domain.Accommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	log.Println("Inserted with id: ", result.InsertedID)

	return nil
}

func (store *AccommodationMongoDBStore) Update(accommodation domain.Accommodation) error {
	filter := bson.D{{Key: "_id", Value: accommodation.Id}}

	_, err := store.accommodations.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: accommodation}})

	return err
}

func (store *AccommodationMongoDBStore) Delete(id primitive.ObjectID) error {
	var deleteAcc domain.Accommodation
	filter := bson.D{{Key: "_id", Value: id}}
	err := store.accommodations.FindOneAndDelete(context.TODO(), filter).Decode(&deleteAcc)

	return err
}


func (repo *AccommodationMongoDBStore) DeleteAllByHost(host string) error {
	filter := bson.D{{Key: "host", Value: host}}
	_, err := repo.accommodations.DeleteMany(context.Background(), filter)
	return err
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var Accommodation domain.Accommodation
		err = cursor.Decode(&Accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &Accommodation)
	}
	err = cursor.Err()
	return
}
