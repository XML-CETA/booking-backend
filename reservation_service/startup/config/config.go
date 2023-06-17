package config

import "os"

type Config struct {
	Port            string
	ReservationDBHost string
	ReservationDBPort string
	AuthServiceHost string
	AuthServicePort string
	AccommodationServiceHost string
	AccommodationServicePort string
  NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
  ProminentHostSubject string
  NotificationSubject string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
		AuthServiceHost: os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort: os.Getenv("AUTH_SERVICE_PORT"),
		AccommodationServiceHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationServicePort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
    NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
    ProminentHostSubject: os.Getenv("PROMINENT_HOST_SUBJECT"),
    NotificationSubject: os.Getenv("NOTIFICATION_SUBJECT"),
	}
}
