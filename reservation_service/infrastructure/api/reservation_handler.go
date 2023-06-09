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
	user, err := Authorize(ctx, []string{"REGULAR"})
	if err != nil {
		return nil, err
	}

	newReservation := domain.MakeReservation(request.Guests, request.Accommodation, user, request.DateFrom, request.DateTo)

	err = h.service.CreateReservation(newReservation)
	if err != nil {
		return nil, err
	}

	return &pb.ReservationCreateResponse{
		Data: fmt.Sprintf("Created"),
	}, nil
}

func (h ReservationHandler) IsAppointmentReserved(ctx context.Context, request *pb.IsAppointmentReservedRequest) (*pb.IsAppointmentReservedResponse, error) {
	reserved, err := h.service.IsAppointmentReserved(request.Accommodation, request.DateFrom, request.DateTo)
	if err != nil {
		return nil, err
	}
	return &pb.IsAppointmentReservedResponse{
		Reserved: reserved,
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

func (h ReservationHandler) GetWaitingReservations(ctx context.Context, request *pb.WaitingReservationsForHostRequest) (*pb.WaitingReservationsForHostResponse, error) {
	host, err := Authorize(ctx, []string{"HOST"})
	if err != nil {
		return nil, err
	}
	waitingReservations, err := h.service.GetWaitingReservations(host)
	return &pb.WaitingReservationsForHostResponse{
		Reservations: h.service.ConvertToGrpcWaitingReservations(waitingReservations),
	}, err
}

func (h ReservationHandler) ConfirmReservation(ctx context.Context, request *pb.ConfirmReservationRequest) (*pb.ConfirmReservationResponse, error) {
	_, err := Authorize(ctx, []string{"HOST"})
	if err != nil {
		return nil, err
	}

	err = h.service.ConfirmReservation(request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmReservationResponse{
		Message: "Reservation is confirmed",
	}, nil
}

func (h ReservationHandler) Delete(ctx context.Context, request *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	user, err := Authorize(ctx, []string{"REGULAR"})
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

func (h ReservationHandler) Decline(ctx context.Context, request *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	_, err := Authorize(ctx, []string{"HOST"})
	if err != nil {
		return nil, err
	}

	err = h.service.Decline(request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteReservationResponse{
		Message: "Successfully declined",
	}, nil
}

func (h ReservationHandler) GetHostAnalytics(ctx context.Context, request *pb.HostAnalyticsRequest) (*pb.HostAnalyticsResponse, error) {
	cancelRate, err := h.service.GetCancelRate(request.Host)
	if err != nil {
		return nil, err
	}

	expiredCount, err := h.service.GetExpiredCount(request.Host)
	if err != nil {
		return nil, err
	}

	intervalCount, err := h.service.GetIntervalCount(request.Host)
	if err != nil {
		return nil, err
	}

	return &pb.HostAnalyticsResponse{
		CancelRate:    cancelRate,
		ExpiredCount:  expiredCount,
		IntervalCount: intervalCount,
	}, nil
}

func (h ReservationHandler) HasLeftoverReservations(ctx context.Context, request *pb.LeftoverReservationsRequest) (*pb.LeftoverReservationsResponse, error) {
	hasActiveReservations, err := h.service.HasActiveReservations(request.User, request.Role)

	if err != nil {
		return nil, err
	}

	return &pb.LeftoverReservationsResponse{
		CanDelete: !hasActiveReservations,
	}, nil
}

func Authorize(ctx context.Context, roleGuard []string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

	if err != nil {
		return "", err
	}

	return user.UserEmail, nil
}
