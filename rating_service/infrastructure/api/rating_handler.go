package api

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	pb "booking-backend/common/proto/rating_service"
	"booking-backend/rating-service/application"
	"booking-backend/rating-service/domain"
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

	newRate := domain.MakeRate(request.Rate, request.Accommodation, user)

	err = h.service.CreateAccommodationRate(newRate)
	if err != nil {
		return nil, err
	}

	return &pb.RateResponse{
		Data: fmt.Sprintf("Created"),
	}, nil
}

func (h RatingHandler) CreateUserRating(ctx context.Context, request *pb.RateUserRequest) (*pb.RateResponse, error) {
	rating := domain.MakeRating(request)
	err := h.service.CreateUserRate(&rating)
	if err != nil {
		return &pb.RateResponse{
			Data: fmt.Sprintf(err.Error()),
		}, err
	}
	return &pb.RateResponse{
		Data: fmt.Sprintf("Created rating."),
	}, nil
}

func (h RatingHandler) UpdateUserRating(ctx context.Context, request *pb.UpdateUserRatingRequest) (*pb.RateResponse, error) {
	err := h.service.UpdateUserRating(request.RatedUser, request.RatedBy, request.Rate)
	if err != nil {
		return &pb.RateResponse{
			Data: fmt.Sprintf(err.Error()),
		}, err
	}
	return &pb.RateResponse{
		Data: fmt.Sprintf("Rating changed"),
	}, nil
}

func (h RatingHandler) UpdateAccommodationRate(ctx context.Context, request *pb.RatingAccommodationRequest) (*pb.RateResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	updateRate := domain.MakeRate(request.Rate, request.Accommodation, user)

	err = h.service.UpdateAccommodationRate(updateRate)
	if err != nil {
		return nil, err
	}

	return &pb.RateResponse{
		Data: fmt.Sprintf("Updated"),
	}, nil
}

func (h RatingHandler) DeleteAccommodationRate(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.RateResponse, error) {
	user, err := Authorize(ctx, "REGULAR")
	if err != nil {
		return nil, err
	}

	err = h.service.DeleteAccommodationRate(user, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.RateResponse{
		Data: fmt.Sprintf("Deleted"),
	}, nil
}

func (h RatingHandler) GetAllAccommodationRates(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.AllAccommodationRatesResponse, error) {
	rates, err := h.service.GetAllAccommodationRates(request.Id)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (h RatingHandler) GetHostRates(ctx context.Context, request *pb.HostRatesRequest) (*pb.HostRatesResponse, error) {
	rates, err := h.service.GetHostRates(request.Id)
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func (h RatingHandler) GetAverageAccommodationRate(ctx context.Context, request *pb.RateAccommodationIdRequest) (*pb.AverageRateAccommodationResponse, error) {
	avgRate, err := h.service.GetAverageAccommodationRate(request.Id)
	if err != nil {
		return nil, err
	}
	return avgRate, nil
}

func Authorize(ctx context.Context, roleGuard string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

	return user.UserEmail, err
}
