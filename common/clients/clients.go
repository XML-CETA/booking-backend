package clients

import (
	"booking-backend/common/proto/accommodation_service"
	"booking-backend/common/proto/auth_service"
	"booking-backend/common/proto/reservation_service"
	users_service "booking-backend/common/proto/user_service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccommodationClient(address string) accommodation_service.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation_service.NewAccommodationServiceClient(conn)
}

func NewUsersClient(address string) users_service.UsersServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return users_service.NewUsersServiceClient(conn)
}

func NewAuthClient(address string) auth_service.AuthServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Auth service: %v", err)
	}
	return auth_service.NewAuthServiceClient(conn)
}

func NewReservationClient(address string) reservation_service.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservation_service.NewReservationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
