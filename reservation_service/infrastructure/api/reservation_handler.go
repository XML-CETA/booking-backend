package api

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/application"
	"booking-backend/reservation-service/domain"
	"booking-backend/reservation-service/startup/config"
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
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
	user , err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	newReservation := domain.MakeReservation(request.Guests, request.Accommodation, user, request.DateFrom, request.DateTo)

	err = h.service.CreateReservation(newReservation)
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

func (h ReservationHandler) Delete(ctx context.Context, request *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	err = h.service.Delete(request.Id, user)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteReservationResponse{
		Message: "Successfully deleted",
	}, nil
}


func Authorize(ctx context.Context, roleGuard string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user , err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

	return user.UserEmail, err
}
