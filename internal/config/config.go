package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST           string
	DB_PORT           string
	DB_NAME           string
	DB_USER           string
	DB_PASSWORD       string
	MIGRATIONS_SOURCE string
}

func LoadConfig(filenames ...string) (*Config, error) {
	if err := godotenv.Load(filenames...); err != nil {
		return nil, err
	}
	return &Config{
		DB_HOST:           getStringEnv("DB_HOST", ""),
		DB_PORT:           getStringEnv("DB_PORT", ""),
		DB_NAME:           getStringEnv("DB_NAME", ""),
		DB_USER:           getStringEnv("DB_USER", ""),
		DB_PASSWORD:       getStringEnv("DB_PASSWORD", ""),
		MIGRATIONS_SOURCE: getStringEnv("MIGRATIONS_SOURCE", "file://migrations"),
	}, nil
}

func (c *Config) GetDBConnString() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		c.DB_HOST,
		c.DB_PORT,
		c.DB_NAME,
		c.DB_USER,
		c.DB_PASSWORD,
	)
}

func getIntEnv(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}

func getStringEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	valueStr = strings.ToUpper(valueStr)
	switch valueStr {
	case "TRUE", "YES", "1":
		return true
	case "FALSE", "NO", "0":
		return false
	default:
		return defaultValue
	}
}
