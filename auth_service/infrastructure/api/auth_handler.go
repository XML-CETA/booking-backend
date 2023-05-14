package api

import (
	"booking-backend/auth_service/application"
)

type AuthHandler struct {
	//pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

