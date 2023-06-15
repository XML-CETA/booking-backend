package application

import (
	"booking-backend/rating-service/domain"
)

type RatingService struct {
	rateAccommodationStore domain.RatingAccommodationStore
}

func NewRatingService(accommodationStore domain.RatingAccommodationStore) *RatingService {
	return &RatingService{
		rateAccommodationStore: accommodationStore,
	}
}
