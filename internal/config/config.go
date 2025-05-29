package config

import (
	"fmt"
	"os"

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
		DB_HOST:           getEnv("DB_HOST", "localhost"),
		DB_PORT:           getEnv("DB_PORT", "5432"),
		DB_NAME:           getEnv("DB_NAME", "person_db"),
		DB_USER:           getEnv("DB_USER", "postgres"),
		DB_PASSWORD:       getEnv("DB_PASSWORD", "password"),
		MIGRATIONS_SOURCE: getEnv("MIGRATIONS_SOURCE", "file://storage/migrations"),
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

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
