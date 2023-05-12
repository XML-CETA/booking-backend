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
		return &reservation.ReservationCreateResponse{
			Greeting: fmt.Sprintf("Failed to create: %s", err),
		}, err
	}

	return &reservation.ReservationCreateResponse{
		Greeting: fmt.Sprintf("%+v created", newReservation),
	}, nil
}

