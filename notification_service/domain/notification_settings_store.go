package domain

import "booking-backend/common/proto/notification_service"


type NotificationSettingsStore interface {
  NewUserSettings(settings NotificationSettings) error
  GetUserSettings(user string) (NotificationSettings, error)
  Update(user string, body *notification_service.UpdateUserSettingsRequest) error
}
