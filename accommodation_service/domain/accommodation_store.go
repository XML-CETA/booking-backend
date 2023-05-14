package domain

type AccommodationStore interface {
	GetAll() ([]Accommodation, error)
	Create(accommodation Accommodation) error
}
