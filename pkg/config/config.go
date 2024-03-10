package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppPort string
	AppName string

	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),

		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MySQLPort:     os.Getenv("MYSQL_PORT"),
		MySQLUser:     os.Getenv("MYSQL_USER"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),
		MySQLDatabase: os.Getenv("MYSQL_DATABASE"),
	}

	return cfg
}
