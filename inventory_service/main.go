package main

import (
	"booking-backend/inventory_service/startup"
	cfg "booking-backend/inventory_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
