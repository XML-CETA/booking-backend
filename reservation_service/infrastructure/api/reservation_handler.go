package api

import (
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/application"
	"booking-backend/reservation-service/domain"
	"context"
	"fmt"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service *application.ReservationService
}

func NewReservationHandler(service *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (h ReservationHandler) Create(ctx context.Context, request *pb.ReservationCreateRequest) (*pb.ReservationCreateResponse, error) {

	newReservation := domain.MakeReservation(request.Accommodation, request.Guests, request.Offer, request.DateFrom, request.DateTo)

	err := h.service.CreateReservation(newReservation)
	if err != nil {
		return nil, err;
	}

	return &pb.ReservationCreateResponse{
		Data: fmt.Sprintf("Created"),
	}, nil
}

func (h ReservationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	reservations, err := h.service.GetAll()

	if err != nil {
		return nil, err
	}

	return &pb.GetAllResponse{
		Reservations: h.service.ConvertToGrpcList(reservations),
	}, nil
}
