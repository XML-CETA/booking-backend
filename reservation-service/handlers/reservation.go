package handlers

import (
	"context"
	"example/grpc/proto/reservation"
	"example/grpc/utils"
	"fmt"
)

type ReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
}

func (h ReservationHandler) Create(ctx context.Context, request *reservation.ReservationCreateRequest) (*reservation.ReservationCreateResponse, error) {

	_, cancel := utils.GetDbClient()
	defer cancel()

	return &reservation.ReservationCreateResponse{
		Greeting: fmt.Sprintf("Hi RESERVATIONS!"),
	}, nil
}

