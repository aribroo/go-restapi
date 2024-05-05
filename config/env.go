package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost string
	AppPort string
	DBUser  string
	DBPass  string
	DBAddr  string
	DBName  string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		AppHost: getEnv("APP_HOST", "http://localhost"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBUser:  getEnv("DB_USER", "root"),
		DBPass:  getEnv("DB_PASSWORD", ""),
		DBAddr:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:  getEnv("DB_NAME", "go_ecommerce"),
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}
