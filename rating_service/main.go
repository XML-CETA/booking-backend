package main

import (
	"booking-backend/rating-service/startup"
	"booking-backend/rating-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
