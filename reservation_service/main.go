package main

import (
	"booking-backend/reservation-service/startup"
	"booking-backend/reservation-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
