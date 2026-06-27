package platform

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPAddr string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	URL             string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MinConnLifetime time.Duration
}

func Load() Config {
	return Config{
		HTTPAddr: getString("HTTP_ADDR", ":8080"),
		Database: DatabaseConfig{
			URL:             getString("DATABASE_URL", "postgres://xbrl:xbrl@localhost:5432/xbrl_validator?sslmode=disable"),
			MaxConns:        int32(getInt("MAX_CONNS", 20)),
			MinConns:        int32(getInt("MIN_CONNS", 2)),
			MaxConnLifetime: getDuration("MAX_CONN_LIFETIME", time.Hour),
			MinConnLifetime: getDuration("MIN_CONN_LIFETIME", 30*time.Minute),
		},
	}
}

func getString(param string, fallback string) string {
	value := os.Getenv(param)
	if value == "" {
		return fallback
	}
	return value
}

func getInt(param string, fallback int) int {
	value := os.Getenv(param)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func getDuration(param string, fallback time.Duration) time.Duration {
	value := os.Getenv(param)
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}
