package domain

type RatingAccommodationStore interface {
	Create(rate RatingAccommodation) error
	Update(rate RatingAccommodation) error
	Delete(rate RatingAccommodation) error
	GetByUserAndAccommodation(userEmail, accommodation string) (RatingAccommodation, error)
	GetAllByAccommodation(accommodation string) ([]RatingAccommodation, error)
}
