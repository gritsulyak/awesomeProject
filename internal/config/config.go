package config

import "fmt"

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
	Addr string
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
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		d.Hosts, d.Port, d.User, d.Name, d.Password, d.SSLMode, d.Schema,
	)
}

func NewDefaultConfig() *Config {
	return &Config{
		HttpServerConfig: HttpServerConfig{
			ListenAddress:  "0.0.0.0:10080",
			MetricsAddress: "0.0.0.0:6060",
		},
		Redis: Redis{
			Addr: "localhost:6379",
		},
		Database: Database{
			Hosts:    "localhost",
			Port:     5432,
			User:     "postgres",
			Name:     "postgres",
			Password: "postgres",
			SSLMode:  "disable",
			Schema:   "public",
		},
	}
}
