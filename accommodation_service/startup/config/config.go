package config

import "os"

type Config struct {
	Port                   string
	AccommodationDBHost    string
	AccommodationDBPort    string
	AuthServiceHost        string
	AuthServicePort        string
	ReservationServiceHost string
	ReservationServicePort string
	UserServiceHost        string
	UserServicePort        string
}

func NewConfig() *Config {
	return &Config{
		Port:                   os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:    os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:    os.Getenv("ACCOMMODATION_DB_PORT"),
		AuthServiceHost:        os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort:        os.Getenv("AUTH_SERVICE_PORT"),
		ReservationServiceHost: os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationServicePort: os.Getenv("RESERVATION_SERVICE_PORT"),
		UserServiceHost:        os.Getenv("USER_SERVICE_HOST"),
		UserServicePort:        os.Getenv("USER_SERVICE_PORT"),
	}
}
