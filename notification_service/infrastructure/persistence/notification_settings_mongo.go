package persistence

import (
	"booking-backend/notification_service/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notifications"
	COLLECTION = "notificationSettings"
)

type NotificationSettingsDB struct {
	notificationSettings *mongo.Collection
}

func NewNotificationSettingsDB(client *mongo.Client) domain.NotificationSettingsStore {
	notifications := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationSettingsDB{
		notificationSettings: notifications,
	}
}
