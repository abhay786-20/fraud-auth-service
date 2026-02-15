package config

import (
	"time"

	"github.com/abhay786-20/fraud-auth-service/pkg/constants"
	"github.com/abhay786-20/fraud-auth-service/pkg/env"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Host    string
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

type AuthConfig struct {
	JWTSecret string
	TokenTTL  time.Duration
}

func LoadConfig(environment *env.Environment) *Config {
	return &Config{
		Server: ServerConfig{
			Host:    environment.Get(constants.EnvServerHost, "0.0.0.0"),
			Port:    environment.Get(constants.EnvServerPort, "8081"),
			GinMode: environment.Get(constants.EnvGinMode, "debug"),
		},
		Database: DatabaseConfig{
			Host:         environment.Get(constants.EnvDBHost, "localhost"),
			Port:         environment.Get(constants.EnvDBPort, "5432"),
			User:         environment.Get(constants.EnvDBUser),
			Password:     environment.Get(constants.EnvDBPassword),
			Name:         environment.Get(constants.EnvDBName),
			MaxOpenConns: environment.GetInt(constants.EnvDBMaxOpenConns, 25),
			MaxIdleConns: environment.GetInt(constants.EnvDBMaxIdleConns, 25),
			MaxLifetime:  time.Duration(environment.GetInt(constants.EnvDBMaxLifetimeMin, 5)) * time.Minute,
		},
		Auth: AuthConfig{
			JWTSecret: environment.Get(constants.EnvJWTSecret),
			TokenTTL:  time.Duration(environment.GetInt(constants.EnvJWTTTLHours, 24)) * time.Hour,
		},
	}
}
