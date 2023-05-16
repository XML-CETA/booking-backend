package application

import (
	"booking-backend/auth_service/infrastructure/services"
	"booking-backend/common/proto/auth_service"
	users_service "booking-backend/common/proto/user_service"
	"context"
	"fmt"
	"time"

	"booking-backend/auth_service/startup/config"

	"github.com/golang-jwt/jwt/v5"
)
type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{
	}
}

func (service *AuthService) Login(ctx context.Context, request *auth_service.AuthenticateRequest) (string, error) {
	userServiceClient := services.NewUsersClient(fmt.Sprintf("%s:%s", config.NewConfig().UserServiceHost, config.NewConfig().UserServicePort))

	login, err := userServiceClient.LoginCheck(ctx, &users_service.LoginRequest{
		Email: request.Username,
		Password: request.Password,
	})

	if err != nil {
		return "", err
	}

	return service.GenerateJwt(login.GetUser())
}

func (service *AuthService) GenerateJwt(user *users_service.User) (string, error) {

	secretKey := []byte("hahabratemoj")

	claims := config.Claims{
		CustomClaims: map[string]string{
			"email": user.Email,
			"role":     user.Role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(secretKey)

	return token, err
}
