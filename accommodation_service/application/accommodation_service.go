package application

import (
	"booking-backend/accommodation_service/domain"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) error {
	return service.store.Create(accommodation)
}
