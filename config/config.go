package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config is used to store all configuration data.
type Config struct {
	Host     string
	Port     string
	From     string
	Password string
}

// GetConfig is used to get all configuration data.
func GetConfig() *Config {
	// Create config struct
	config := &Config{}

	err := godotenv.Load("config/config.env")
	if err != nil {
		panic(err)
	}

	// Set configuration data
	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
	config.From = os.Getenv("FROM")
	config.Password = os.Getenv("PASSWORD")

	// Return config struct
	return config
}
