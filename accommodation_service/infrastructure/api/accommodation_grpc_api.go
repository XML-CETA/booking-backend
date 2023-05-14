package api

import (
	"context"
	"fmt"

	"booking-backend/accommodation_service/application"
	"booking-backend/accommodation_service/domain"

	pb "booking-backend/common/proto/accommodation_service"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllAccommodationRequest) (*pb.GetAllAccommodationResponse, error) {
	accommodations, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	return &pb.GetAllAccommodationResponse{
		Accommodations: handler.service.ConvertToGrpcList(accommodations),
	}, nil
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *pb.AccommodationCreateRequest) (*pb.AccommodationCreateResponse, error) {

	newAccommodation := domain.MakeAccommodation(request.Longitude, request.Latitude, request.MinGuests, request.MaxGuests, request.Name)

	fmt.Println(newAccommodation)
	err := handler.service.Create(newAccommodation)
	if err != nil {
		return nil, err
	}

	return &pb.AccommodationCreateResponse{
		Data: fmt.Sprintf("Successfully Created!"),
	}, nil
}
