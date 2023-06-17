package persistence

import (
	"booking-backend/notification_service/domain"

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
