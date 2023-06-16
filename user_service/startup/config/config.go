package config

import "os"

type Config struct {
	Port        string
	UsersDBHost string
	UsersDBPort string
	AuthServiceHost string
	AuthServicePort string
  NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
  ProminentHostSubject string
}

func NewConfig() *Config {
	return &Config{
		Port:        os.Getenv("USER_SERVICE_PORT"),
		UsersDBHost: os.Getenv("USERS_DB_HOST"),
		UsersDBPort: os.Getenv("USERS_DB_PORT"),
		AuthServiceHost: os.Getenv("AUTH_SERVICE_HOST"),
		AuthServicePort: os.Getenv("AUTH_SERVICE_PORT"),
    NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
    ProminentHostSubject: os.Getenv("PROMINENT_HOST_SUBJECT"),
	}
}
