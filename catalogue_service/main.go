package main

import (
	"booking-backend/catalogue_service/startup"
	cfg "booking-backend/catalogue_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
