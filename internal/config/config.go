package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Port        string
	Environment string
	JWTSecret   string
	CORSOrigins []string
	Database    *DatabaseConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host                   string
	Port                   string
	User                   string
	Password               string
	DBName                 string
	SSLMode                string
	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetimeMinutes int
}

// Load loads configuration from environment and flags
func Load() (*Config, error) {
	// Load .env file (if exists)
	_ = godotenv.Load()

	cfg := &Config{
		Port:        getEnv("PORT", "3000"),
		Environment: getEnv("APP_ENV", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "your-super-secret-key-change-in-production"),
		CORSOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		Database: &DatabaseConfig{
			Host:                   getEnv("DB_HOST", "localhost"),
			Port:                   getEnv("DB_PORT", "5432"),
			User:                   getEnv("DB_USER", "dev"),
			Password:               getEnv("DB_PASSWORD", "dev"),
			DBName:                 getEnv("DB_NAME", "affiliate_db"),
			SSLMode:                getEnv("DB_SSLMODE", "disable"),
			MaxOpenConns:           getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:           getEnvInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetimeMinutes: getEnvInt("DB_CONN_MAX_LIFETIME_MINUTES", 5),
		},
	}

	// Override with flags (optional)
	flagPort := flag.String("port", "", "Server port")
	flag.Parse()
	if *flagPort != "" {
		cfg.Port = *flagPort
	}

	return cfg, nil
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction returns true if running in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
