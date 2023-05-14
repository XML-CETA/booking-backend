package startup

import (
<<<<<<< HEAD
	cfg "booking-backend/api_gateway/startup/config"
	accommodationGw "booking-backend/common/proto/accommodation_service"
=======
	"booking-backend/common/proto/auth_service"
>>>>>>> 573ca28 (Add jwt library)
	"booking-backend/common/proto/reservation_service"
	users_service "booking-backend/common/proto/user_service"
	"context"
	"fmt"
	"log"
	"net/http"

<<<<<<< HEAD
=======
	cfg "booking-backend/api_gateway/startup/config"

>>>>>>> 573ca28 (Add jwt library)
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
	log.Print(reservationEndpoint)
	err := reservation_service.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = users_service.RegisterUsersServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
=======
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	log.Print(authEndpoint)
	err = auth_service.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
>>>>>>> 573ca28 (Add jwt library)
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}
=======

>>>>>>> 573ca28 (Add jwt library)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
