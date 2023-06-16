package domain

type RatingUserStore interface {
	Create(rating RatingUser) error
}
