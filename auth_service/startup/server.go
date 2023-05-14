package startup

import (
	"booking-backend/auth_service/application"
	"booking-backend/auth_service/infrastructure/api"
	"booking-backend/auth_service/startup/config"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	reservationService := server.initAuthService()

	reservationHandler := server.initAuthHandler(reservationService)

	server.startGrpcServer(reservationHandler)
}


func (server *Server) initAuthService() *application.AuthService {
	return application.NewAuthService()
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(productHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//reservationProto.RegisterAuthServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
