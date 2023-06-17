package application

import "booking-backend/notification_service/domain"

type NotificationService struct {
	notificationSettings domain.NotificationSettingsStore
	notifications domain.NotificationStore
}

func NewNotificationService(storeSettings domain.NotificationSettingsStore, notifications domain.NotificationStore) *NotificationService {
	return &NotificationService{
		notificationSettings: storeSettings,
    notifications: notifications,
	}
}
