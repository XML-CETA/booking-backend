package main

import (
	"booking-backend/notification_service/startup/config"
	"booking-backend/notification_service/startup"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
