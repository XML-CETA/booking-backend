package api

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	pb "booking-backend/common/proto/rating_service"
	"booking-backend/rating-service/application"
	"booking-backend/rating-service/startup/config"
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type RatingHandler struct {
	pb.UnimplementedRatingServiceServer
	service *application.RatingService
}

func NewRatingHandler(service *application.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (h RatingHandler) CreateAccommodationRate(ctx context.Context, request *pb.RatingAccommodationRequest) (*pb.RateResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	// err = h.service.CreateReservation(newReservation)
	// if err != nil {
	// 	return nil, err
	// }

	// return &pb.ReservationCreateResponse{
	// 	Data: fmt.Sprintf("Created"),
	// }, nil
	fmt.Println("Created")
	return nil, nil
}

func (h RatingHandler) UpdateAccommodationRate(ctx context.Context, request *pb.RatingAccommodationRequest) (*pb.RateResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	fmt.Println("Update")
	return nil, nil
}

func (h RatingHandler) DeleteAccommodationRate(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.RateResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	fmt.Println("Delete")
	return nil, nil
}

func (h RatingHandler) GetAllAccommodationRates(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.AllAccommodationRatesResponse, error) {
	fmt.Println("GetAll")
	return nil, nil
}

func (h RatingHandler) GetAverageAccommodationRate(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.AverageRateAccommodationResponse, error) {
	fmt.Println("GetAverage")
	return nil, nil
}

func Authorize(ctx context.Context, roleGuard string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

	return user.UserEmail, err
}
