package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	// 1. Load .env file (if exists)
	_ = godotenv.Load() // Ignore error if file not found

	// 2. Define default values
	cfg := &Config{
		Port: "3000",
	}

	// 3. Override with Environment Variables
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	// 4. Override with Flags (Optional, but good for testing/scripts)
	flagPort := flag.String("port", "", "Server port")
	flag.Parse()
	if *flagPort != "" {
		cfg.Port = *flagPort
	}

	return cfg, nil
}
