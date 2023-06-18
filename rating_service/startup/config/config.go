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
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	RateUserCommandSubject   string
	RateUserReplySubject     string
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
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		RateUserCommandSubject:   os.Getenv("RATE_USER_COMMAND_SUBJECT"),
		RateUserReplySubject:     os.Getenv("RATE_USER_REPLY_SUBJECT"),
	}
}
