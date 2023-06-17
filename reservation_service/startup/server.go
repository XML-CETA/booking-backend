package startup

import (
	"booking-backend/common/messaging"
	reservationProto "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/application"
	"booking-backend/reservation-service/domain"
	"booking-backend/reservation-service/infrastructure/api"
	"booking-backend/reservation-service/infrastructure/persistence"
	"booking-backend/reservation-service/startup/config"
	"fmt"
	"log"
	"net"

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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()

	reservationStore := server.initReservationStore(mongoClient)

  publisher := server.initPublisher(server.config.ProminentHostSubject)
  notificationPublisher := server.initPublisher(server.config.NotificationSubject)

	reservationService := server.initReservationService(reservationStore, publisher, notificationPublisher)

	reservationHandler := server.initReservationHandler(reservationService)

	server.startGrpcServer(reservationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationStore(client *mongo.Client) domain.ReservationStore {
	store := persistence.NewReservationMongoDBStore(client)
	return store
}

func (server *Server) initReservationService(store domain.ReservationStore, prominentHostPublisher messaging.PublisherModel, notificationPublisher messaging.PublisherModel) *application.ReservationService {
	return application.NewReservationService(store, prominentHostPublisher, notificationPublisher)
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}

func (server *Server) startGrpcServer(productHandler *api.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservationProto.RegisterReservationServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initPublisher(subject string) messaging.PublisherModel {
	publisher, err := messaging.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}
