package startup

import (
	"fmt"
	"log"
	"net"

	"booking-backend/accommodation_service/application"
	"booking-backend/accommodation_service/domain"
	"booking-backend/accommodation_service/infrastructure/api"
	"booking-backend/accommodation_service/infrastructure/persistance"
	"booking-backend/accommodation_service/startup/config"

	accommodation "booking-backend/common/proto/accommodation_service"

	"go.mongodb.org/mongo-driver/mongo"
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

const (
	QueueGroup = "accommodation_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accommodationStore := server.initAccommodationStore(mongoClient)

	accommodationService := server.initAccommodationService(accommodationStore)

	accommodationHandler := server.initAccommodationHandler(accommodationService)

	server.startGrpcServer(accommodationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistance.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationStore(client *mongo.Client) domain.AccommodationStore {
	store := persistance.NewAccommodationMongoDBStore(client)
	return store
}

func (server *Server) initAccommodationService(store domain.AccommodationStore) *application.AccommodationService {
	return application.NewAccommodationService(store)
}

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
