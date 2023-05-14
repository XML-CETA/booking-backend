package config

import "os"

type Config struct {
	Port          string
	ReservationHost string
	ReservationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("GATEWAY_PORT"),
		ReservationHost: os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort: os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
