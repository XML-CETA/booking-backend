package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccommodationStore interface {
	GetAll() ([]Accommodation, error)
	Create(accommodation Accommodation) error
	Update(accommodation Accommodation) error
	Delete(id primitive.ObjectID) error
	GetById(id primitive.ObjectID) (*Accommodation, error)
  DeleteAllByHost(host string) error
}
