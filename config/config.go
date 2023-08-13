package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_Name     string
	DB_User     string
	DB_Password string
	DB_Host     string
	DB_Port     int
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dbPortStr := os.Getenv("DB_PORT")
	dbPortInt, err := strconv.Atoi(dbPortStr)

	if err != nil {
		log.Println("Failed to convert port to a number")
		return nil
	}

	config := &Config{
		DB_Name:     dbName,
		DB_User:     dbUser,
		DB_Password: dbPassword,
		DB_Host:     dbHost,
		DB_Port:     dbPortInt,
	}

	return config
}
