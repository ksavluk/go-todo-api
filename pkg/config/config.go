package config

import (
	"os"
)

type Config struct {
	Address string
	DSN     string
}

func LoadFromEnv() Config {
	return Config{
		Address: os.Getenv("ADDRESS"),
		DSN:     os.Getenv("DSN"),
	}
}
