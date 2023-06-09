package config

import "os"

type Config struct {
	Port              string
	UserHost          string
	UserPort          string
	ReservationHost   string
	ReservationPort   string
	AccommodationHost string
	AccommodationPort string
	RatingHost        string
	RatingPort        string
	AuthHost          string
	AuthPort          string
  NotificationHost string
  NotificationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		ReservationHost:   os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		RatingHost:        os.Getenv("RATING_SERVICE_HOST"),
		RatingPort:        os.Getenv("RATING_SERVICE_PORT"),
		AuthHost:          os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:          os.Getenv("AUTH_SERVICE_PORT"),
    NotificationHost: os.Getenv("NOTIFICATION_SERVICE_HOST"),
    NotificationPort: os.Getenv("NOTIFICATION_SERVICE_PORT"),
	}
}
