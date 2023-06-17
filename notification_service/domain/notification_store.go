package domain

type NotificationStore interface {
  Create(notification Notification) error
  GetAllByUser(user string) ([]Notification, error)
}
