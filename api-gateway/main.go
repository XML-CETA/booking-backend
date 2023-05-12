package main

import (
	"context"
	"example/gateway/config"
	"example/gateway/proto/greeter"
	"example/gateway/proto/reservation"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()

	greeterConn, err := grpc.DialContext(
		context.Background(),
		cfg.GreeterServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	reservationConn, err := grpc.DialContext(
		context.Background(),
		cfg.ReservationServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	greeterClient := greeter.NewGreeterServiceClient(greeterConn)
	err = greeter.RegisterGreeterServiceHandlerClient(
		context.Background(),
		gwmux,
		greeterClient,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Register Reservation
	reservationClient := reservation.NewReservationServiceClient(reservationConn)
	err = reservation.RegisterReservationServiceHandlerClient(
		context.Background(),
		gwmux,
		reservationClient,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
