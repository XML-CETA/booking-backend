package persistence

import (
	"booking-backend/notification_service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (store *NotificationSettingsDB) NewUserSettings(settings domain.NotificationSettings) error {
	_, err := store.notificationSettings.InsertOne(context.Background(), settings)

  return err
}

func (store *NotificationSettingsDB) GetUserSettings(user string) (domain.NotificationSettings, error) {
	filter := bson.D{
		{Key: "user", Value: user},
	}
	var result domain.NotificationSettings
	err := store.notificationSettings.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
