package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingUserStore interface {
	Create(rating *RatingUser) (primitive.ObjectID, error)
	UpdateStatus(host, user string, status Status) error
	GetHostRates(host string) ([]RatingUser, error)
	GetByHostAndUser(host, user string, id primitive.ObjectID) (RatingUser, error)
}
