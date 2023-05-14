package main

import (
	"booking-backend/ordering_service/startup"
	cfg "booking-backend/ordering_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
