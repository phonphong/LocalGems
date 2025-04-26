package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	"localgems/internal/interfaces/api/controllers/handlers"
	"localgems/internal/interfaces/api/controllers/middlewares"
)

func SetupRouter(coffeeHandler *handlers.CoffeeHandler) *gin.Engine {

	router := gin.New()

	router.Use(middlewares.Recovery())
	router.Use(middlewares.CORS())
	router.Use(middlewares.ErrorHandler())

	rateLimiter := middlewares.NewRateLimiter(100, time.Minute)
	router.Use(rateLimiter.RateLimit())

	// API routes
	api := router.Group("/api/v1")
	{

		coffees := api.Group("/coffees")
		{
			// Public endpoints
			coffees.GET("", coffeeHandler.GetAllCoffees)
			coffees.GET("/:id", coffeeHandler.GetCoffeeByID)
			coffees.GET("/search", coffeeHandler.SearchCoffees)

			// Protected endpoints
			authRoutes := coffees.Group("/")
			authRoutes.Use(middlewares.Auth()) // Áp dụng Auth middleware chỉ cho routes này
			{
				authRoutes.POST("", coffeeHandler.CreateCoffee)
				authRoutes.PUT("/:id", coffeeHandler.UpdateCoffee)
				authRoutes.DELETE("/:id", coffeeHandler.DeleteCoffee)
			}
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}
