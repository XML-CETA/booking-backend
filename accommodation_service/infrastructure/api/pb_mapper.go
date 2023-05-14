package api

import (
	"booking-backend/accommodation_service/domain"

	pb "booking-backend/common/proto/accommodation_service"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:        accommodation.Id.Hex(),
		Longitude: accommodation.Longitude,
		Latitude:  accommodation.Latitude,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Name:      accommodation.Name,
	}
	return accommodationPb
}
