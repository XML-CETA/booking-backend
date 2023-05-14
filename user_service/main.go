package main

import (
	"booking-backend/user-service/startup"
	"booking-backend/user-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
