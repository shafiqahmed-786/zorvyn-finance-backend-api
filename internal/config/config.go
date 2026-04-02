package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port      string
    DBHost    string
    DBPort    string
    DBUser    string
    DBPass    string
    DBName    string
    JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))

    return &Config{
        Port:      os.Getenv("PORT"),
        DBHost:    os.Getenv("DB_HOST"),
        DBPort:    os.Getenv("DB_PORT"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASSWORD"),
        DBName:    os.Getenv("DB_NAME"),
        JWTSecret: os.Getenv("JWT_SECRET"),
    }
}