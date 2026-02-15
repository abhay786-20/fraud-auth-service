package constants

// Environment variable constants for the Fraud Auth Service.
// All environment variables are prefixed with AUTH_ for namespace isolation.

// Server configuration environment variables
const (
	EnvServerHost = "AUTH_SERVER_HOST" // Server bind address (default: "0.0.0.0")
	EnvServerPort = "AUTH_SERVER_PORT" // Server port (default: "8081")
	EnvGinMode    = "GIN_MODE"         // Gin mode: debug, release, test (default: "debug")
)

// Database configuration environment variables
const (
	EnvDBHost           = "AUTH_DB_HOST"             // PostgreSQL host (default: "localhost")
	EnvDBPort           = "AUTH_DB_PORT"             // PostgreSQL port (default: "5432")
	EnvDBUser           = "AUTH_DB_USER"             // PostgreSQL user (REQUIRED)
	EnvDBPassword       = "AUTH_DB_PASSWORD"         // PostgreSQL password (REQUIRED)
	EnvDBName           = "AUTH_DB_NAME"             // PostgreSQL database name (REQUIRED)
	EnvDBMaxOpenConns   = "AUTH_DB_MAX_OPEN_CONNS"   // Max open connections (default: 25)
	EnvDBMaxIdleConns   = "AUTH_DB_MAX_IDLE_CONNS"   // Max idle connections (default: 25)
	EnvDBMaxLifetimeMin = "AUTH_DB_MAX_LIFETIME_MIN" // Connection max lifetime in minutes (default: 5)
)

// Authentication configuration environment variables
const (
	EnvJWTSecret  = "AUTH_JWT_SECRET"    // JWT signing secret (REQUIRED)
	EnvJWTTTLHours = "AUTH_JWT_TTL_HOURS" // JWT token TTL in hours (default: 24)
)

// RequiredEnvVars contains all environment variables that MUST be set.
// Application will fail to start if any of these are missing.
var RequiredEnvVars = []string{
	EnvDBUser,
	EnvDBPassword,
	EnvDBName,
	EnvJWTSecret,
}

// OptionalEnvVars contains environment variables with default values.
// These don't need to be explicitly set.
var OptionalEnvVars = []string{
	EnvServerHost,
	EnvServerPort,
	EnvDBHost,
	EnvDBPort,
	EnvDBMaxOpenConns,
	EnvDBMaxIdleConns,
	EnvDBMaxLifetimeMin,
	EnvJWTTTLHours,
}
