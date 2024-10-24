package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                        string
	DB_Host                     string
	Port                        string
	DBPort                      int64
	User                        string
	Password                    string
	DBName                      string
	SSLMode                     string
	JWTTokenExpirationInSeconds int64
	JWTSecret                   string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:                        getEnv("HOST", "http://localhost"),
		DB_Host:                     getEnv("DB_HOST", "http://localhost"),
		User:                        getEnv("DB_USER", "root"),
		Password:                    getEnv("DB_PASSWORD", "root"),
		SSLMode:                     getEnv("SSSMODE", "disable"),
		Port:                        getEnv("PORT", "5050"),
		DBPort:                      getEnvAsInt("DB_PORT", 5432),
		DBName:                      getEnv("DB_NAME", "hotelRoom"),
		JWTTokenExpirationInSeconds: getEnvAsInt("JWTTokenExpirationInSeconds", 3600*24),
		JWTSecret:                   getEnv("JWT_SECRET", "secret"),
	}
}

// getEnv retrieves the value of the environment variable or uses a fallback if it's not set.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt retrieves the value of the environment variable as an int64 or returns the fallback.
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return intValue
		}
	}
	return fallback
}
