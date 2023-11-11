package config

import (
	"os"
	"otpapi/model"

	"github.com/joho/godotenv"
)

func GetConfig() *model.Config {
	// Create config struct
	config := &model.Config{}

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
