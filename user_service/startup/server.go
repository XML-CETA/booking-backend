package startup

import (
	"fmt"
	"log"
	"net"

	"booking-backend/common/messaging"
	usersProto "booking-backend/common/proto/user_service"
	"booking-backend/user-service/application"
	"booking-backend/user-service/domain"
	"booking-backend/user-service/infrastructure/api"
	"booking-backend/user-service/infrastructure/persistence"
	"booking-backend/user-service/startup/config"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const (
	QueueGroup = "user_service"
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

	userStore := server.initUserStore(mongoClient)

	subscriber := server.initSubscriber(server.config.ProminentHostSubject, QueueGroup)
	publisher := server.initPublisher(server.config.NotificationSubject)

	userService := server.initUserService(userStore, subscriber, publisher)

	userHandler := server.initUserHandler(userService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UsersDBHost, server.config.UsersDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	return store
}

func (server *Server) initUserService(store domain.UserStore, subscriber messaging.SubscriberModel, publisher messaging.PublisherModel) *application.UserService {
	service, err := application.NewUserService(store, subscriber, publisher)
	if err != nil {
		log.Fatalf("Failed to start service %v", err)
	}

	return service
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	usersProto.RegisterUsersServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initSubscriber(subject, queueGroup string) messaging.SubscriberModel {
	subscriber, err := messaging.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
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
