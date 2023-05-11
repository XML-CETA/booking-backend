package handlers

import (
	"context"
	"example/grpc/proto/reservation"
	"fmt"
)

type ReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
}

func (h ReservationHandler) Create(ctx context.Context, request *reservation.ReservationCreateRequest) (*reservation.ReservationCreateResponse, error) {
	return &reservation.ReservationCreateResponse{
		Greeting: fmt.Sprintf("Hi RESERVATIONS!"),
	}, nil
}

