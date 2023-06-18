package startup

import (
	"booking-backend/common/messaging"
	pb "booking-backend/common/proto/rating_service"
	"booking-backend/rating-service/application"
	"booking-backend/rating-service/domain"
	"booking-backend/rating-service/infrastructure/api"
	"booking-backend/rating-service/infrastructure/persistence"
	"booking-backend/rating-service/startup/config"
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

const (
	QueueGroup = "rate_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	ratingAccommodationStore := server.initRatingAccommodationStore(mongoClient)
	ratingUserStore := server.initRatingUserStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.RateUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.RateUserReplySubject, QueueGroup)
	rateUserOrchestrator := server.initRateUserOrchestrator(commandPublisher, replySubscriber)

	ratingService := server.initRatingService(ratingAccommodationStore, ratingUserStore, rateUserOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.RateUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RateUserReplySubject)

	server.initRateUserHandler(ratingService, replyPublisher, commandSubscriber)

	ratingHandler := server.initRatingHandler(ratingService)

	server.startGrpcServer(ratingHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.RatingDBHost, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRatingAccommodationStore(client *mongo.Client) domain.RatingAccommodationStore {
	store := persistence.NewRatingAccommodationMongoDBStore(client)
	return store
}

func (server *Server) initRatingUserStore(client *mongo.Client) domain.RatingUserStore {
	store := persistence.NewRatingUserMongoDBStore(client)
	return store
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

func (server *Server) initSubscriber(subject, queueGroup string) messaging.SubscriberModel {
	subscriber, err := messaging.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initRateUserOrchestrator(publisher messaging.PublisherModel, subscriber messaging.SubscriberModel) *application.RateUserOrchestrator {
	orchestrator, err := application.NewRateUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initRateUserHandler(service *application.RatingService, publisher messaging.PublisherModel, subscriber messaging.SubscriberModel) {
	_, err := api.NewRateUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initRatingService(rateAccommodationStore domain.RatingAccommodationStore, rateUserStore domain.RatingUserStore, orchestrator *application.RateUserOrchestrator) *application.RatingService {
	return application.NewRatingService(rateAccommodationStore, rateUserStore, orchestrator)
}

func (server *Server) initRatingHandler(service *application.RatingService) *api.RatingHandler {
	return api.NewRatingHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *api.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRatingServiceServer(grpcServer, ratingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
