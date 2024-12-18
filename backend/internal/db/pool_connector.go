package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"quote-management/internal/config"
)

// Connect initializes the PostgreSQL connection pool
func Connect(cfg config.DatabaseSettings) (*pgxpool.Pool, error) {
	dbConnString := getConnectionString(&cfg)
	var pool *pgxpool.Pool
	var err error

	// Retry logic
	for i := 0; i < 3; i++ {
		pool, err = pgxpool.Connect(context.Background(), dbConnString)
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database. Retrying... (%d/3)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	fmt.Println("Database connection established")
	return pool, nil
}

// getConnectionString formats the connection string from the config
func getConnectionString(cfg *config.DatabaseSettings) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

}
