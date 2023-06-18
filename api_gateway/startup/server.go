package startup

import (
	cfg "booking-backend/api_gateway/startup/config"
	accommodationGw "booking-backend/common/proto/accommodation_service"
	"booking-backend/common/proto/auth_service"
	"booking-backend/common/proto/notification_service"
	"booking-backend/common/proto/rating_service"
	"booking-backend/common/proto/reservation_service"
	users_service "booking-backend/common/proto/user_service"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err := reservation_service.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = users_service.RegisterUsersServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err = auth_service.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingHost, server.config.RatingPort)
	err = rating_service.RegisterRatingServiceHandlerFromEndpoint(context.TODO(), server.mux, ratingEndpoint, opts)
	if err != nil {
		panic(err)
	}

	notificationEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationHost, server.config.NotificationPort)
	err = notification_service.RegisterNotificationServiceHandlerFromEndpoint(context.TODO(), server.mux, notificationEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(server.mux)))
}

func cors(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
    if r.Method == "OPTIONS" {
     return
    }
    h.ServeHTTP(w, r)
  })
}
