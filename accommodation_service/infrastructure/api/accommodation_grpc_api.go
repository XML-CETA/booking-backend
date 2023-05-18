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
	acc, err := domain.MakeAccommodation(request)
	if err != nil {
		return nil, err
	}

	err = handler.service.Update(acc)
	if err != nil {
		return &pb.Response{
			Data: fmt.Sprint(err.Error()),
		}, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully updated"),
	}, nil
}

func (handler *AccommodationHandler) Delete(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.Response, error) {
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

func (handler *AccommodationHandler) CreateAppointment(ctx context.Context, request *pb.SingleAppointment) (*pb.Response, error) {
	createAppointment, err := domain.MakeCreateAppointment(request)
	if err != nil {
		return nil, err
	}

	err = handler.service.AddAppointment(createAppointment)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Data: fmt.Sprintf("Succesfully created!"),
	}, nil
}

func (handler *AccommodationHandler) UpdateAppointment(ctx context.Context, request *pb.UpdateAppointmentRequest) (*pb.Response, error) {
	updateAppointment, err := domain.MakeUpdateAppointment(request)

	if err != nil {
		return nil, err
	}

	err = handler.service.UpdateAppointment(updateAppointment)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully updated!"),
	}, nil
}
