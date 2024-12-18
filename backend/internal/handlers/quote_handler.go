package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"quote-management/internal/models"
)

// CreateQuote handles POST requests to create a new quote.
func (h *Handler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote
	// Debug: Log incoming request body
	log.Println("Reading request body...")

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		log.Printf("Decoding error: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received quote: %+v", quote)

	// Call the repository to create the quote
	createdQuote, err := h.Repo.CreateQuote(r.Context(), &quote)
	if err != nil {
		http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		return
	}

	// Return the created quote as a JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdQuote)
}
