package api

import (
	pb "booking-backend/common/proto/rating_service"
	"booking-backend/rating-service/application"
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
