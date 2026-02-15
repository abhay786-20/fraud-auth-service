package config

import (
	"log"
	"os"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func Load() *Config {

	cfg := &Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	validate(cfg)

	return cfg
}

func validate(cfg *Config) {
	if cfg.Port == "" ||
		cfg.DBHost == "" ||
		cfg.DBUser == "" ||
		cfg.DBName == "" ||
		cfg.JWTSecret == "" {
		log.Fatal("Missing required environment variables")
	}
}