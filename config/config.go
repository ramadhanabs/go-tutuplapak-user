package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpiryHours int
}

// LoadConfig loads configuration from environment variables or a `.env` file.
func LoadConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	viper.AutomaticEnv()

	config := Config{
		DBHost:         viper.GetString("DB_HOST"),
		DBPort:         viper.GetString("DB_PORT"),
		DBUser:         viper.GetString("DB_USER"),
		DBPassword:     viper.GetString("DB_PASSWORD"),
		DBName:         viper.GetString("DB_NAME"),
		JWTSecret:      viper.GetString("JWT_SECRET"),
		JWTExpiryHours: viper.GetInt("JWT_EXPIRY_HOURS"),
	}

	if config.JWTExpiryHours == 0 {
		config.JWTExpiryHours = 24 // Default to 24 hours
	}

	return config
}
