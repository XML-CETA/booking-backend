package main

import (
	"booking-backend/accommodation_service/startup"
	cfg "booking-backend/accommodation_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
