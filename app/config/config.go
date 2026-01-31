package config

import (
	"log"
	"os"
)

type DBConfig struct {
	DatabaseURL string
}

func LoadDBConfig() *DBConfig {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		log.Fatal("DATABASE_URL is required")
	}

	return &DBConfig{
		DatabaseURL: dbURL,
	}
}
