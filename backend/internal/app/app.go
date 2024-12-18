package app

import (
	"context"
	"log"
	"quote-management/internal/config"
	"quote-management/internal/db"
	"quote-management/internal/handlers"
	"quote-management/internal/repository"
	"quote-management/internal/server"
	"time"
)

type App struct {
	cfg    *config.Config
	server *server.Server
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Run(ctx context.Context) error {
	// 1. Initialize Database Connection
	dbConn, err := db.Connect(app.cfg.DatabaseSettings)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	// Initialize Repository with DB Pool
	repo := repository.NewRepository(dbConn)

	// Initialize Handlers
	handler := handlers.NewHandler(repo)

	// 4. Create Server
	app.server = server.NewServer(app.cfg.AppSettings.ServerPort, handler)

	// 5. Start Server
	app.server.Start()

	log.Println("Application is running...")
	return nil
}

func (app *App) StopServer() error {
	// Graceful Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := app.server.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("Server stopped successfully")
	return nil
}
