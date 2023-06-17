package config

import "os"

type Config struct {
	Port            string
	NotificationDBHost string
	NotificationDBPort string
	AuthServiceHost string
	AuthServicePort string
  NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
  NotificationSubject string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("NOTIFICATION_SERVICE_PORT"),
		NotificationDBHost: os.Getenv("NOTIFICATION_DB_HOST"),
		NotificationDBPort: os.Getenv("NOTIFICATION_DB_PORT"),
		AuthServiceHost: os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort: os.Getenv("AUTH_SERVICE_PORT"),
    NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
    NotificationSubject: os.Getenv("NOTIFICATION_SUBJECT"),
	}
}
