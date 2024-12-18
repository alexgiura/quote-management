package test

import (
	"context"
	"quote-management/internal/models"
)

type MockRepository struct {
	CreateQuoteFunc    func(ctx context.Context, quote *models.Quote) (*models.Quote, error)
	CreateCustomerFunc func(ctx context.Context, customer *models.Customer) (*models.Customer, error)
}

func (m *MockRepository) CreateQuote(ctx context.Context, quote *models.Quote) (*models.Quote, error) {
	return m.CreateQuoteFunc(ctx, quote)
}

func (m *MockRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if m.CreateCustomerFunc != nil {
		return m.CreateCustomerFunc(ctx, customer)
	}
	return nil, nil
}
