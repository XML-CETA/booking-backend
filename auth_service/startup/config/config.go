package config

import "os"

type Config struct {
	Port string
	UserServiceHost string
	UserServicePort string
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("AUTH_SERVICE_PORT"),
		UserServiceHost: os.Getenv("USER_SERVICE_HOST"),
		UserServicePort: os.Getenv("USER_SERVICE_PORT"),
	}
}
