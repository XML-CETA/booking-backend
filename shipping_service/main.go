package main

import (
	"booking-backend/shipping_service/startup"
	cfg "booking-backend/shipping_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
