package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type NotificationSettings struct{
  Id primitive.ObjectID `bson:"_id, omitempty"`
  User string
  ReservationRequest bool
  ReservationCancel bool
  PersonalRating bool
  AccommodationRating bool
  ProminentStatusChange bool
  ReservationResponse bool
}

type Notification struct {
  Id primitive.ObjectID `bson:"_id, omitempty"`
  User string
  Subject string
  Content string
  Viewed bool
}
