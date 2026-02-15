package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/abhay786-20/fraud-auth-service/pkg/constants"
	"github.com/joho/godotenv"
)

type Environment struct{}

func New() (*Environment, error) {
	log.Println("[INFO] Loading environment variables from .env file")

	if err := godotenv.Load(); err != nil {
		log.Printf("[ERROR] Failed to load .env file: %s", err)
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	log.Println("[INFO] Environment variables loaded successfully")

	// Validate required environment variables
	if err := validateRequired(constants.RequiredEnvVars); err != nil {
		return nil, err
	}

	log.Println("[INFO] Required environment variables validated")
	return &Environment{}, nil
}

// validateRequired checks if all required environment variables are set.
func validateRequired(requiredVars []string) error {
	var missing []string

	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		log.Printf("[ERROR] Missing required environment variables: %s", strings.Join(missing, ", "))
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return nil
}

func (e *Environment) Get(key string, defaultVal ...string) string {
	val := os.Getenv(key)
	if val == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return val
}

func (e *Environment) GetInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

func (e *Environment) GetBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return b
}
