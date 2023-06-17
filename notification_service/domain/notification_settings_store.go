package domain


type NotificationSettingsStore interface {
  NewUserSettings(settings NotificationSettings) error
  GetUserSettings(user string) (NotificationSettings, error)
}
