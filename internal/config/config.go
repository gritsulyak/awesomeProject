package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HttpServerConfig HttpServerConfig
	Database         Database
	Redis            Redis
}

type HttpServerConfig struct {
	ListenAddress  string
	MetricsAddress string
}

type Redis struct {
	Addr     string
	Password string
}

type Database struct {
	Name     string
	Schema   string
	Hosts    string
	User     string
	Password string
	Port     int
	SSLMode  string
}

func (d *Database) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		d.Hosts, d.Port, d.User, d.Name, d.Password, d.SSLMode, d.Schema,
	)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getenvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func NewConfigFromEnv() *Config {
	return &Config{
		HttpServerConfig: HttpServerConfig{
			ListenAddress:  getenv("HTTP_LISTEN", "0.0.0.0:10080"),
			MetricsAddress: getenv("METRICS_LISTEN", "0.0.0.0:6080"),
		},
		Redis: Redis{
			Addr:     getenv("REDIS_ADDR", "redis:7379"),
			Password: getenv("REDIS_PASSWORD", ""),
		},
		Database: Database{
			Hosts:    getenv("DB_HOST", "postgres"),
			Port:     getenvInt("DB_PORT", 5432),
			User:     getenv("DB_USER", "postgres"),
			Name:     getenv("DB_NAME", "postgres"),
			Password: getenv("DB_PASSWORD", "postgres"),
			SSLMode:  getenv("DB_SSLMODE", "disable"),
			Schema:   getenv("DB_SCHEMA", "public"),
		},
	}
}
