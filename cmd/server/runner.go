package main

import (
	"localgems/config"
	"localgems/internal/app/usecases"
	"localgems/internal/infra/db"
	"localgems/internal/infra/repositories"
	"localgems/internal/interfaces/api/controllers/handlers"
	"localgems/internal/interfaces/api/routes"
	"log"
)

func main() {
	// Load config
	cfg := config.NewConfig()

	// Initialize database
	db, err := db.NewSQLiteConnection(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Setup repositories
	coffeeRepo := repositories.NewSQLiteCafeRepository(db)

	// Setup usecases
	coffeeUsecase := usecases.NewCoffeeUsecase(coffeeRepo)

	// Setup handlers
	coffeeHandler := handlers.NewCoffeeHandler(coffeeUsecase)

	// Setup router
	router := routes.SetupRouter(coffeeHandler)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
