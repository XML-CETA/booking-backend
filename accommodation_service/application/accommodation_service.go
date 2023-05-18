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

func (service *AccommodationService) GetAll() ([]*pb.SingleAccommodation, error) {
	accommodations, err := service.store.GetAll()
	grpcAccommodations := service.ConvertToGrpcList(accommodations)
	return grpcAccommodations, err
}

func (service *AccommodationService) GetById(id string) (*pb.SingleAccommodation, error) {
	accommodation, err := service.store.GetById(id)
	grpcAccommodation := service.ConvertToGrpc(accommodation)
	return grpcAccommodation, err
}

func (service *AccommodationService) Create(accommodation domain.Accommodation) error {
	return service.store.Create(accommodation)
}

func (service *AccommodationService) Update(accommodation domain.Accommodation) error {
	return service.store.Update(accommodation)
}

func (service *AccommodationService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *AccommodationService) ConvertToGrpcList(accommodations []domain.Accommodation) []*pb.SingleAccommodation {
	var converted []*pb.SingleAccommodation

	for _, entity := range accommodations {
		newRes := service.ConvertToGrpc(&entity)
		converted = append(converted, newRes)
	}

	return converted
}

func (service *AccommodationService) ConvertToGrpc(accommodation *domain.Accommodation) *pb.SingleAccommodation {

	res := pb.SingleAccommodation{
		Id:        accommodation.Id.Hex(),
		Longitude: accommodation.Longitude,
		Latitude:  accommodation.Latitude,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Name:      accommodation.Name,
		Address: &pb.AccommodationAddress{
			City:    accommodation.Address.City,
			Number:  accommodation.Address.Number,
			Country: accommodation.Address.Country,
			Street:  accommodation.Address.Street,
		},
	}

	return &res
}
