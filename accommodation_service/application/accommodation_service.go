package application

import (
	"booking-backend/accommodation_service/domain"
	pb "booking-backend/common/proto/accommodation_service"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) GetAll() ([]domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(accommodation domain.Accommodation) error {
	return service.store.Create(accommodation)
}

func (service *AccommodationService) ConvertToGrpcList(accommodations []domain.Accommodation) []*pb.SingleAccommodation {
	var converted []*pb.SingleAccommodation

	for _, entity := range accommodations {
		newRes := pb.SingleAccommodation{
			Id:        entity.Id.Hex(),
			Longitude: entity.Longitude,
			Latitude:  entity.Latitude,
			MinGuests: entity.MinGuests,
			MaxGuests: entity.MaxGuests,
			Name:      entity.Name,
		}

		converted = append(converted, &newRes)
	}

	return converted
}
