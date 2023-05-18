package domain

type AccommodationStore interface {
	GetAll() ([]Accommodation, error)
	Create(accommodation Accommodation) error
	Update(accommodation Accommodation) error
	Delete(id string) error
	GetById(id string) (*Accommodation, error)
}
