package startup

import (
	"booking-backend/common/proto/notification_service"
	"booking-backend/notification_service/application"
	"booking-backend/notification_service/domain"
	"booking-backend/notification_service/infrastructure/api"
	"booking-backend/notification_service/infrastructure/persistence"
	"booking-backend/notification_service/startup/config"
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

  notificationSettings := server.initNotificationSettingsStore(mongoClient)
  notifications := server.initNotificationStore(mongoClient)

  notificationService := server.initNotificationService(notificationSettings, notifications)

  notificationHandler := server.initNotificationHandler(notificationService)

  server.startGrpcServer(notificationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationDBHost, server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationSettingsStore(client *mongo.Client) domain.NotificationSettingsStore {
	store := persistence.NewNotificationSettingsDB(client)
	return store
}

func (server *Server) initNotificationStore(client *mongo.Client) domain.NotificationStore {
	store := persistence.NewNotificationDB(client)
	return store
}

func (server *Server) initNotificationService(storeSettings domain.NotificationSettingsStore, notifications domain.NotificationStore) *application.NotificationService {
	return application.NewNotificationService(storeSettings, notifications)
}

func (server *Server) initNotificationHandler(service *application.NotificationService) *api.NotificationHandler {
	return api.NewNotificationHandler(service)
}

func (server *Server) startGrpcServer(handler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification_service.RegisterNotificationServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
