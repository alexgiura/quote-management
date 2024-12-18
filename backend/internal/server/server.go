package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"quote-management/internal/handlers"
)

type Server struct {
	Router *mux.Router
	HTTP   *http.Server
}

func NewServer(port string, handler *handlers.Handler) *Server {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/api/quotes", handler.CreateQuote).Methods("POST")
	//router.HandleFunc("/api/quotes", handler.GetQuotes).Methods("GET")
	//router.HandleFunc("/api/quotes/{id}", handler.GetQuote).Methods("GET")
	//router.HandleFunc("/api/quotes/{id}", handler.UpdateQuote).Methods("PUT")
	//router.HandleFunc("/api/quotes/{id}", handler.DeleteQuote).Methods("DELETE")

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	return &Server{
		Router: router,
		HTTP:   httpServer,
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.HTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("Server failed to start: %s", err))
		}
	}()
	fmt.Println("Server started on", s.HTTP.Addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HTTP.Shutdown(ctx)
}
