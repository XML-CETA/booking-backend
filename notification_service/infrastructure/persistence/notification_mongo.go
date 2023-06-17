package persistence

import (
	"booking-backend/notification_service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASENOTIFICATION   = "notifications"
	COLLECTIONNOTIFICATION = "notifications"
)

type NotificationDB struct {
	notifications *mongo.Collection
}

func NewNotificationDB(client *mongo.Client) domain.NotificationStore {
	notifications := client.Database(DATABASENOTIFICATION).Collection(COLLECTIONNOTIFICATION)
	return &NotificationDB{
		notifications: notifications,
	}
}

func (store *NotificationDB) Create(notification domain.Notification) error {
	_, err := store.notifications.InsertOne(context.Background(), notification)
  return err
}

func (store *NotificationDB) GetAllByUser(user string) ([]domain.Notification, error) {
	filter := bson.D{{ Key: "user", Value: user }}
	var notifications []domain.Notification
	dataResult, err := store.notifications.Find(context.Background(), filter)
	for dataResult.Next(context.TODO()) {
		var notification domain.Notification
		err := dataResult.Decode(&notification)
		if err == nil {
			notifications = append(notifications, notification)
		}
	}

	return notifications, err
}
