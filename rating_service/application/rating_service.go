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

func (service *RatingService) GetRateByUserAndAccommodation(userEmail string, accommodationId string) (domain.RatingAccommodation, error) {
	rate, err := service.rateAccommodationStore.GetById(objId)
	return grpcAccommodation, err
}
