package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"quote-management/internal/handlers"
	"quote-management/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuoteHandlerWithJSON(t *testing.T) {
	// Load test data from JSON file
	mockQuote := &models.Quote{}
	LoadJSON(t, "TestData/test_quote.json", mockQuote)

	// Set up the mock repository
	mockRepo := &MockRepository{
		CreateQuoteFunc: func(ctx context.Context, quote *models.Quote) (*models.Quote, error) {
			quote.ID = 1 // Simulate DB insertion
			return quote, nil
		},
	}

	// Initialize the real handler
	handler := handlers.NewHandler(mockRepo)

	// Convert the loaded quote struct to JSON
	payload, _ := json.Marshal(mockQuote)
	req := httptest.NewRequest("POST", "/api/quotes", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rr := httptest.NewRecorder()

	// Call the actual HTTP handler
	handler.CreateQuote(rr, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Parse the response
	var createdQuote models.Quote
	err := json.NewDecoder(rr.Body).Decode(&createdQuote)
	assert.NoError(t, err)

	// Verify the response fields
	assert.Equal(t, "Q123456", createdQuote.QuoteCode)
	assert.Equal(t, 1, createdQuote.CustomerID)
	assert.Equal(t, "PENDING", createdQuote.Status)
	assert.Equal(t, 2500.00, createdQuote.TotalAmount)
}
