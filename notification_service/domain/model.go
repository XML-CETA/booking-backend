package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type NotificationSettings struct{
  Id primitive.ObjectID `bson:"_id,omitempty"`
  User string
  Role string
  ReservationRequest bool
  ReservationCancel bool
  PersonalRating bool
  AccommodationRating bool
  ProminentStatusChange bool
  ReservationResponse bool
}

type Notification struct {
  Id primitive.ObjectID `bson:"_id,omitempty"`
  User string
  Subject string
  Content string
  Viewed bool
}

func MakeSettingsFromRole(user, role string) NotificationSettings {
  isHost := false

  if role == "HOST" {
    isHost = true
  }

  return NotificationSettings{
    User: user,
    Role: role,
    ReservationRequest: isHost,
    ReservationCancel: isHost,
    PersonalRating: isHost,
    AccommodationRating: isHost,
    ProminentStatusChange: isHost,
    ReservationResponse: !isHost,
  }
}
