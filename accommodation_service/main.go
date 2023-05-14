package main

import (
	"booking-backend/accommodation_service/startup"
	cfg "booking-backend/accommodation_service/startup/config"
	"log"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Print(config)
	server.Start()
}
