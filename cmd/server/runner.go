package main

import (
	"local-gems-server/config"
	"local-gems-server/internal/app/usecases"
	"local-gems-server/internal/infra/db"
	"local-gems-server/internal/infra/repositories"
	"local-gems-server/internal/interfaces/api/controllers/handlers"
	"local-gems-server/internal/interfaces/api/routes"
	"log"
)

func main() {
	cfg := config.NewConfig()

	database, err := db.NewMySQLConnection(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to MySQL: %v", err)
	}
	defer database.Close()

	localRepo := repositories.NewMySQLLocalRepository(database)
	localUsecase := usecases.NewLocalUsecase(localRepo)
	localHandler := handlers.NewLocalHandler(localUsecase)

	router := routes.SetupRouter(localHandler)

	log.Printf("Server running at %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
