package config

import "os"

type Config struct {
	Address               string
	GreeterServiceAddress string
	ReservationServiceAddress string
}

func GetConfig() Config {
	return Config{
		GreeterServiceAddress: os.Getenv("GREETER_SERVICE_ADDRESS"),
		ReservationServiceAddress: os.Getenv("RESERVATION_SERVICE_ADDRESS"),
		Address:               os.Getenv("GATEWAY_ADDRESS"),
	}
}
