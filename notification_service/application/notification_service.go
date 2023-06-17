package application

import (
	"booking-backend/common/proto/notification_service"
	"booking-backend/notification_service/domain"
	"errors"
)

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

func (service *NotificationService) NewUserSettings(host string, role string) error {
  _, err := service.notificationSettings.GetUserSettings(host)
  if err == nil {
    return errors.New("User settings already exist")
  }

  return service.notificationSettings.NewUserSettings(domain.MakeSettingsFromRole(host, role))
}

func (service *NotificationService) GetUserSettings(host string) (domain.NotificationSettings, error) {
  return service.notificationSettings.GetUserSettings(host)
}

func (service *NotificationService) UpdateUserSettings(host string, body *notification_service.UpdateUserSettingsRequest) (domain.NotificationSettings, error) {
  err := service.notificationSettings.Update(host, body)
  if err != nil {
    return domain.NotificationSettings{}, err
  }

  return service.GetUserSettings(host)
}
