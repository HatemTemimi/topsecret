package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all the environment variables
type Config struct {
	Env                string
	GooglePlacesAPIKey string
	MainPort           int
	DatabaseName       string
	DatabaseUser       string
	DatabasePassword   string
	DatabasePort       int
	DatabaseHost       string
}

// LoadConfig reads the environment variables and populates the Config struct
func LoadConfig() (*Config, error) {
	// Determine the environment and load the corresponding .env file
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}
	envFile := "../config/.env." + env

	// Load the environment variables from the file
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Error loading %s file, falling back to default .env", envFile)
		godotenv.Load(".env") // Optional fallback to a default .env
	}

	// Parse environment variables into Config struct
	config := &Config{
		Env:                GetEnv("GO_ENV", "dev"),
		GooglePlacesAPIKey: GetEnv("GOOGLE_PLACES_API_KEY", ""),
		MainPort:           GetEnvAsInt("MAIN_PORT", 3001),
		DatabaseName:       GetEnv("DATABASE_NAME", "database"),
		DatabaseUser:       GetEnv("DATABASE_USER", "user"),
		DatabasePassword:   GetEnv("DATABASE_PASSWORD", "secret"),
		DatabasePort:       GetEnvAsInt("DATABASE_PORT", 27017),
		DatabaseHost:       GetEnv("DATABASE_HOST", "localhost"),
	}

	return config, nil
}

// Helper function to read an environment variable or fallback to a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to read an integer environment variable or fallback to a default value
func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
