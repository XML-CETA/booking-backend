package handlers

import (
	"context"
	"example/grpc/model"
	"example/grpc/proto/reservation"
	"example/grpc/service"
	"fmt"
)

type ReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
	Service *service.Service
}

func (h ReservationHandler) Create(ctx context.Context, request *reservation.ReservationCreateRequest) (*reservation.ReservationCreateResponse, error) {

	newReservation := model.MakeReservation(request.Accommodation, request.Guests, request.Offer, request.DateFrom, request.DateTo)

	err := h.Service.CreateReservation(newReservation)
	if err != nil {
		return nil, err;
	}

	return &reservation.ReservationCreateResponse{
		Data: fmt.Sprintf("Created"),
	}, nil
}

func (h ReservationHandler) GetAll(ctx context.Context, request *reservation.GetAllRequest) (*reservation.GetAllResponse, error) {
	reservations, err := h.Service.GetAll()

	if err != nil {
		return nil, err
	}

	return &reservation.GetAllResponse{
		Reservations: h.Service.ConvertToGrpcList(reservations),
	}, nil
}
