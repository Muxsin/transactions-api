package config

import "os"

type Config struct {
	AppName        string
	HTTPServerPort string
}

func New() *Config {
	return &Config{
		AppName:        os.Getenv("APP_NAME"),
		HTTPServerPort: os.Getenv("HTTP_SERVER_PORT"),
	}
}
