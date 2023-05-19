package application

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	users_service "booking-backend/common/proto/user_service"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"booking-backend/auth_service/startup/config"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)
type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{
	}
}

func (service *AuthService) Login(ctx context.Context, request *auth_service.AuthenticateRequest) (string, error) {
	userServiceClient := clients.NewUsersClient(fmt.Sprintf("%s:%s", config.NewConfig().UserServiceHost, config.NewConfig().UserServicePort))

	login, err := userServiceClient.LoginCheck(ctx, &users_service.LoginRequest{
		Email: request.Email,
		Password: request.Password,
	})

	if err != nil {
		return "", err
	}

	return service.generateJwt(login.GetUser())
}

func (service *AuthService) generateJwt(user *users_service.User) (string, error) {
	secretKey := config.NewConfig().SecretKey

	claims := config.Claims{
		CustomClaims: map[string]string{
			"email":	user.Email,
			"role":     user.Role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString([]byte(secretKey))

	return token, err
}

func (service *AuthService) Authorize(ctx context.Context, roleguard string) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	header := md.Get("authorization")

	if len(header) == 0 {
		return "", errors.New("No authorization header")
	}

	bearer := header[0]

	token, claims := service.parseJwt(bearer)
	if !token.Valid {
		return "", errors.New("Invalid token")
	}

	if claims.CustomClaims["role"] != roleguard {
		return "", errors.New("You are unauthorized for this endpoint")
	}

	return claims.CustomClaims["email"], nil
}

func (service *AuthService) parseJwt(authorizationHeader string) (*jwt.Token, *config.Claims) {
	tokenString := strings.TrimSpace(strings.Split(authorizationHeader, "Bearer")[1])
	token, _ := jwt.ParseWithClaims(tokenString, &config.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.NewConfig().SecretKey), nil
	}, jwt.WithLeeway(5*time.Second))

	claims, _ := token.Claims.(*config.Claims)

	return token, claims
}

