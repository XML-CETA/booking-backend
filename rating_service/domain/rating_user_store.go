package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingUserStore interface {
	Create(rating *RatingUser) (primitive.ObjectID, error)
	UpdateStatus(id primitive.ObjectID, status Status) error
	GetHostRates(host string) ([]RatingUser, error)
	GetByHostAndUser(host, user string, id primitive.ObjectID) (RatingUser, error)
	Update(host, user string, rate int32) error
}
