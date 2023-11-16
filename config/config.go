package config

import (
	"log"
	"otpapi/model"

	"github.com/spf13/viper"
)

func GetConfig() *model.Config {
	// Create config struct
	config := &model.Config{}
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .yaml file")
		panic(err)
	}

	// Set configuration data
	config.Host = viper.GetString("email.host")
	config.Port = viper.GetString("email.port")
	config.From = viper.GetString("email.from")
	config.Password = viper.GetString("email.password")
	config.Key = viper.GetString("security.key")

	// Return config struct
	return config
}
