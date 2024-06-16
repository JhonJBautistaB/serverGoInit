package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PortServer string
	NameServer string
}

func LoadConfig() *Config {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		PortServer: os.Getenv("PORT_SERVER"),
		NameServer: os.Getenv("NAME_SERVER"),
	}
}
