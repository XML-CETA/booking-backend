package persistance

import (
	"context"
	"log"

	"booking-backend/accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson"
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

func (store *AccommodationMongoDBStore) Create(accommodation domain.Accommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	log.Println("Inserted with id: ", result.InsertedID)

	return nil
}

// func (store *OrderMongoDBStore) DeleteAll() {
// 	store.orders.DeleteMany(context.TODO(), bson.D{{}})
// }

// func (store *OrderMongoDBStore) UpdateStatus(order *domain.Order) error {
// 	result, err := store.orders.UpdateOne(
// 		context.TODO(),
// 		bson.M{"_id": order.Id},
// 		bson.D{
// 			{"$set", bson.D{{"status", order.Status}}},
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if result.MatchedCount != 1 {
// 		return errors.New("one document should've been updated")
// 	}
// 	return nil
// }

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

// func (store *OrderMongoDBStore) filterOne(filter interface{}) (Order *domain.Order, err error) {
// 	result := store.orders.FindOne(context.TODO(), filter)
// 	err = result.Decode(&Order)
// 	return
// }

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
