package messaging

type NotificationType int32
const (
  ReservationRequest NotificationType = iota
  ReservationCancel
  PersonalRating
  AccommodationRating
  ProminentStatusChange
  ReservationResponse
)

type NotificationMessage struct{
  Type NotificationType
  User string
  Subject string
  Content string
}
