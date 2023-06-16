package config

import "os"

type Config struct {
	Port                     string
	RatingDBHost             string
	RatingDBPort             string
	AuthServiceHost          string
	AuthServicePort          string
	AccommodationServiceHost string
	AccommodationServicePort string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost:             os.Getenv("RATING_DB_HOST"),
		RatingDBPort:             os.Getenv("RATING_DB_PORT"),
		AuthServiceHost:          os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort:          os.Getenv("AUTH_SERVICE_PORT"),
		AccommodationServiceHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationServicePort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
	}
}
