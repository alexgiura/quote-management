package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"quote-management/internal/models"
	"quote-management/internal/sqlc"
)

// Repository manages database queries
type RepositoryInterface interface {
	CreateQuote(ctx context.Context, quote *models.Quote) (*models.Quote, error)
}

// Repository manages database queries
type Repository struct {
	DBProvider *sqlc.Queries
	DBPool     *pgxpool.Pool
}

// NewRepository initializes the real repository
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		DBProvider: sqlc.New(pool),
		DBPool:     pool,
	}
}
