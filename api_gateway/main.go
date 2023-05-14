package main

import (
	"booking-backend/api_gateway/startup"
	"booking-backend/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
