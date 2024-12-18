package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

// AppSettings holds configuration related to the application server.
type AppSettings struct {
	ServerPort  string `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost  string `env:"SERVER_HOST" envDefault:"localhost"`
	GraphQLPath string `env:"GRAPHQL_PATH" envDefault:"/graphql"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	DebugMode   bool   `env:"DEBUG_MODE" envDefault:"false"`
}

// DatabaseSettings holds configuration related to the PostgreSQL database.
type DatabaseSettings struct {
	DbName   string `env:"POSTGRES_DB_NAME" envDefault:"db"`
	Host     string `env:"POSTGRES_DB_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_DB_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_DB_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_DB_PASSWORD" envDefault:"postgres"`
}

// Config holds all the configuration settings for the application.
type Config struct {
	AppSettings      AppSettings
	DatabaseSettings DatabaseSettings
}

// Load loads the configuration from environment variables and returns a Config struct.
func Load() (*Config, error) {
	cfg := &Config{}

	// Load environment variables from .env file (if exists)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables.")
	}

	// Parse the configuration from environment variables
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error loading configuration: %s", err)
	}

	return cfg, nil
}

// DBConnectionString generates the database connection string based on the Config.
func (cfg *Config) DBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DatabaseSettings.User,
		cfg.DatabaseSettings.Password,
		cfg.DatabaseSettings.Host,
		cfg.DatabaseSettings.Port,
		cfg.DatabaseSettings.DbName,
	)
}

// init function loads the .env file when the program starts
func init() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(currentFilePath), "..", "..")
	envFilePath := filepath.Join(rootPath, ".env")

	// Load the .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
