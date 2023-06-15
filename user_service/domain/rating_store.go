package domain

type RatingStore interface {
	Create(rating Rating) error
}
