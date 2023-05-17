package main

import (
	"booking-backend/auth_service/startup"
	"booking-backend/auth_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
