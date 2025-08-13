package loader

import (
	"os"

	"github.com/joho/godotenv"
	logx "github.com/karthikbalasubramani/snap-basket/internal/logger"
	cfg "github.com/karthikbalasubramani/snap-basket/internal/models"
)

// Loggers
var Info = logx.CustomLogger.Info
var Error = logx.CustomLogger.Error
var Warn = logx.CustomLogger.Warn

// Common function to get environment variables
func getEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		Warn.Printf("Unable to load environmenr file for database configuration: %v", err)
	}
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// To load Server Configuration from environment
func LoadGoServerConfig() cfg.ServerConfig {
	cfgServerConfig := cfg.ServerConfig{
		Protocol: getEnv("SERVER_PROTOCOL", "tcp"),
		Port:     getEnv("SERVER_PORT", "8900"),
	}
	return cfgServerConfig
}

// To load Databsae Configuration from environment
func LoadDatabseConfig() cfg.DatabaseConfig {
	cfgDatabaseConfig := cfg.DatabaseConfig{
		DatabaseName:   getEnv("DATABASE_NAME", "snap-basket"),
		Uri:            getEnv("MONGO_URI", "mongodb://localhost:27017"),
		UserCollection: getEnv("DATABASE_COLLECTION_USERS", "users"),
	}
	return cfgDatabaseConfig
}
