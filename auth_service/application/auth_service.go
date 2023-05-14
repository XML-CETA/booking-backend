package application

import (
	pb "booking-backend/common/proto/reservation_service"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{
	}
}

