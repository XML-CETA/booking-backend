package config

import "os"

type Config struct {
	Port                 string
	UsersDBHost          string
	UsersDBPort          string
	AuthServiceHost      string
	AuthServicePort      string
	ReservationHost      string
	ReservationPort      string
	AccommodationHost      string
	AccommodationPort      string
	NotificationHost     string
	NotificationPort     string
	NatsHost             string
	NatsPort             string
	NatsUser             string
	NatsPass             string
	ProminentHostSubject string
	NotificationSubject  string
}

func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("USER_SERVICE_PORT"),
		UsersDBHost:          os.Getenv("USERS_DB_HOST"),
		UsersDBPort:          os.Getenv("USERS_DB_PORT"),
		AuthServiceHost:      os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort:      os.Getenv("AUTH_SERVICE_PORT"),
		ReservationHost:      os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:      os.Getenv("RESERVATION_SERVICE_PORT"),
		AccommodationHost:      os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:      os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		NotificationHost:     os.Getenv("NOTIFICATION_SERVICE_HOST"),
		NotificationPort:     os.Getenv("NOTIFICATION_SERVICE_PORT"),
		NatsHost:             os.Getenv("NATS_HOST"),
		NatsPort:             os.Getenv("NATS_PORT"),
		NatsUser:             os.Getenv("NATS_USER"),
		NatsPass:             os.Getenv("NATS_PASS"),
		ProminentHostSubject: os.Getenv("PROMINENT_HOST_SUBJECT"),
		NotificationSubject:  os.Getenv("NOTIFICATION_SUBJECT"),
	}
}
