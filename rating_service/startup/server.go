package startup

import (
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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	ratingAccommodationStore := server.initRatingAccommodationStore(mongoClient)
	ratingUserStore := server.initRatingUserStore(mongoClient)

	ratingService := server.initRatingService(ratingAccommodationStore, ratingUserStore)

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

func (server *Server) initRatingService(rateAccommodationStore domain.RatingAccommodationStore, rateUserStore domain.RatingUserStore) *application.RatingService {
	return application.NewRatingService(rateAccommodationStore, rateUserStore)
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
