package application

import (
	"booking-backend/common/messaging"
	"booking-backend/common/proto/notification_service"
	"booking-backend/notification_service/domain"
	"errors"
)

type NotificationService struct {
	notificationSettings domain.NotificationSettingsStore
	notifications domain.NotificationStore
  subscriber messaging.SubscriberModel
}

func NewNotificationService(storeSettings domain.NotificationSettingsStore, notifications domain.NotificationStore, subscriber messaging.SubscriberModel) (*NotificationService, error) {
  service := &NotificationService{
		notificationSettings: storeSettings,
    notifications: notifications,
    subscriber: subscriber,
	}

  err := service.subscriber.Subscribe(service.ProcessNotification)

  if err != nil {
    return nil, err
  }

  return service, nil
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

func (service *NotificationService) GetByUser(user string) ([]domain.Notification, error) {
  return service.notifications.GetAllByUser(user)
}

func (service *NotificationService) RedactUser(user string) error {
  err := service.notificationSettings.RedactUser(user)
  if err != nil {
    return err
  }

  err = service.notifications.RedactUser(user)
  return err
}

func (service *NotificationService) ProcessNotification(notification messaging.NotificationMessage) {
  settings, err := service.GetUserSettings(notification.User)
  if err != nil {
    return
  }

  var shouldPush bool

  switch notification.Type {
  case messaging.ReservationRequest:
    shouldPush = settings.ReservationRequest
  case messaging.ReservationCancel:
    shouldPush = settings.ReservationCancel
  case messaging.PersonalRating:
    shouldPush = settings.PersonalRating
  case messaging.AccommodationRating:
    shouldPush = settings.AccommodationRating
  case messaging.ProminentStatusChange:
    shouldPush = settings.ProminentStatusChange
  case messaging.ReservationResponse:
    shouldPush = settings.ReservationResponse
  default:
    shouldPush = false
  }


  if shouldPush {
    service.notifications.Create(domain.Notification{
      User: notification.User,
      Subject: notification.Subject,
      Content: notification.Content,
      Viewed: false,
    })
  }
}
