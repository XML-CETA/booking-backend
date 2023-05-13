package config

import "os"

type Config struct {
	Address                   string
	GreeterServiceAddress     string
	ReservationServiceAddress string
	UserServiceAddress        string
}

func GetConfig() Config {
	return Config{
		GreeterServiceAddress:     os.Getenv("GREETER_SERVICE_ADDRESS"),
		ReservationServiceAddress: os.Getenv("RESERVATION_SERVICE_ADDRESS"),
		Address:                   os.Getenv("GATEWAY_ADDRESS"),
		UserServiceAddress:        os.Getenv("USER_SERVICE_ADDRESS"),
	}
}
