package api

import (
	"context"
	"fmt"
	"log"

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
		Accommodations: accommodations,
	}, nil
}

func (handler *AccommodationHandler) GetById(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.SingleAccommodation, error) {
	accommodation, err := handler.service.GetById(request.Id)

	if err != nil {
		return nil, err
	}

	return accommodation, err
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *pb.AccommodationCreateRequest) (*pb.Response, error) {

	newAccommodation := domain.MakeCreateAccommodation(request)

	err := handler.service.Create(newAccommodation)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Successfully Created!"),
	}, nil
}

func (handler *AccommodationHandler) Update(ctx context.Context, request *pb.SingleAccommodation) (*pb.Response, error) {
	reqAccommodation := request
	acc := domain.MakeAccommodation(reqAccommodation)

	err := handler.service.Update(acc)
	if err != nil {
		return &pb.Response{
			Data: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully updated"),
	}, nil
}

func (handler *AccommodationHandler) Delete(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.Response, error) {
	log.Println(request.Id)
	err := handler.service.Delete(request.Id)
	if err != nil {
		return &pb.Response{
			Data: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully deleted!"),
	}, nil
}
