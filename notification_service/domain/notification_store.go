package domain

type NotificationStore interface {
  Create(notification Notification) error
  RedactUser(user string) error
  GetAllByUser(user string) ([]Notification, error)
}
