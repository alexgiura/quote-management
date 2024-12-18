package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"quote-management/internal/models"
	"quote-management/internal/sqlc"
	"quote-management/internal/util"
)

//

// CreateQuote creates a new quote using SQLC's query
func (r *Repository) CreateQuote(ctx context.Context, quote *models.Quote) (*models.Quote, error) {

	var createdQuoteID int32

	// Start a transaction with BeginFunc

	err := r.DBPool.BeginFunc(ctx, func(tx pgx.Tx) error {

		q := r.DBProvider.WithTx(tx)

		// Create the quote
		created, err := q.CreateQuote(ctx, sqlc.CreateQuoteParams{
			QuoteCode:   util.NullableStr(&quote.QuoteCode),
			CustomerID:  int32(quote.CustomerID),
			Status:      quote.Status,
			TotalAmount: quote.TotalAmount,
			Currency:    quote.Currency,
		})
		if err != nil {
			return err
		}

		// Save the created quote ID
		createdQuoteID = created.ID

		// Save products linked to the created quote
		for _, product := range quote.Products {
			err := q.AddProductToQuote(ctx, sqlc.AddProductToQuoteParams{
				QuoteID:      created.ID,
				ProductID:    int32(product.ProductID),
				Quantity:     int32(product.Quantity),
				PricePerUnit: product.PricePerUnit,
				TaxRate:      product.TaxRate,
				TotalPrice:   product.TotalPrice,
			})
			if err != nil {
				return fmt.Errorf("failed to add product to quote: %v", err)
			}
		}

		return nil
	})

	// If transaction fails, return an error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("transaction failed: %v", err)
	}

	createdQuote, err := r.GetQuoteByID(ctx, int(createdQuoteID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch created quote: %w", err)
	}

	return createdQuote, nil

}

func (r *Repository) GetQuoteByID(ctx context.Context, id int) (*models.Quote, error) {
	// Fetch the quote by ID
	quoteRow, err := r.DBProvider.GetQuoteByID(ctx, int32(id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quote by id: %w", err)
	}

	// Fetch all products linked to the quote
	productsRows, err := r.DBProvider.GetProductsByQuoteID(ctx, int32(id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products for quote: %w", err)
	}

	// Map products to the Go struct
	var products []models.QuoteProduct
	for _, p := range productsRows {
		products = append(products, models.QuoteProduct{
			ProductID:    int(p.ProductID),
			ProductName:  p.ProductName,
			Quantity:     int(p.Quantity),
			PricePerUnit: p.PricePerUnit,
			TaxRate:      p.TaxRate,
			TotalPrice:   p.TotalPrice,
		})
	}

	// Return the complete quote with products
	return &models.Quote{
		ID:          int(quoteRow.QuoteID),
		QuoteCode:   *util.StringOrNil(quoteRow.QuoteCode),
		CustomerID:  int(quoteRow.CustomerID),
		Status:      quoteRow.Status,
		TotalAmount: quoteRow.TotalAmount,
		Currency:    quoteRow.Currency,
		Products:    products,
		CreatedAt:   quoteRow.CreatedAt,
		UpdatedAt:   quoteRow.UpdatedAt,
	}, nil
}
