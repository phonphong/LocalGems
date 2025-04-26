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
	cfg := config.NewConfig()

	database, err := db.NewMySQLConnection(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to MySQL: %v", err)
	}
	defer database.Close()

	coffeeRepo := repositories.NewMySQLCoffeeRepository(database)
	coffeeUsecase := usecases.NewCoffeeUsecase(coffeeRepo)
	coffeeHandler := handlers.NewCoffeeHandler(coffeeUsecase)

	router := routes.SetupRouter(coffeeHandler)

	log.Printf("Server running at %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
