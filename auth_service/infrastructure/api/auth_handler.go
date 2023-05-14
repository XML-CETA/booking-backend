package api

import (
	"booking-backend/auth_service/application"
	pb "booking-backend/common/proto/auth_service"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{
		Token: "",
	}, nil
}

func (h AuthHandler) Authorize(ctx context.Context, request *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {

	return &pb.AuthorizeResponse{
		Success: true,
	}, nil
}


