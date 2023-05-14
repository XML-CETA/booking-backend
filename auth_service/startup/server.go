package startup

import (
	"booking-backend/auth_service/application"
	"booking-backend/auth_service/infrastructure/api"
	"booking-backend/auth_service/startup/config"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	authProto "booking-backend/common/proto/auth_service"
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

	authService := server.initAuthService()

	authHandler := server.initAuthHandler(authService)

	server.startGrpcServer(authHandler)
}


func (server *Server) initAuthService() *application.AuthService {
	return application.NewAuthService()
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authProto.RegisterAuthServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
