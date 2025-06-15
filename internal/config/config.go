package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "clinic"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecret"),
	}
}

func getEnv(key, Default string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return Default
}

func (c *Config) GetDSN() string {
	// fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	return os.Getenv("DATABASE_URL")
}

