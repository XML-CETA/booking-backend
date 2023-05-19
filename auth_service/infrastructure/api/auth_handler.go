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

func (h AuthHandler) Login(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	token, err := h.service.Login(ctx, request)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticateResponse{
		Token: token,
	}, nil
}

func (h AuthHandler) Authorize(ctx context.Context, request *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	email ,err := h.service.Authorize(ctx, request.GetRoleGuard())
	if err != nil {
		return nil, err
	}
	return &pb.AuthorizeResponse{
		Success: true,
		UserEmail: email,
	}, nil
}


