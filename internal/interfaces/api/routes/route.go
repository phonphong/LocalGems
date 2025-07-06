package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	"local-gems-server/internal/interfaces/api/controllers/handlers"
	"local-gems-server/internal/interfaces/api/controllers/middlewares"
)

func SetupRouter(localHandler *handlers.LocalHandler) *gin.Engine {

	router := gin.New()

	router.Use(middlewares.Recovery())
	router.Use(middlewares.CORS())
	router.Use(middlewares.ErrorHandler())

	rateLimiter := middlewares.NewRateLimiter(100, time.Minute)
	router.Use(rateLimiter.RateLimit())

	// API routes
	api := router.Group("/api/v1")
	{

		locals := api.Group("/locals")
		{
			// Public endpoints
			locals.GET("", localHandler.GetAllLocals)
			locals.GET("/:id", localHandler.GetLocalByID)
			locals.GET("/search", localHandler.SearchLocals)

			// Protected endpoints
			authRoutes := locals.Group("/")
			authRoutes.Use(middlewares.Auth()) // Áp dụng Auth middleware chỉ cho routes này
			{
				authRoutes.POST("", localHandler.CreateLocal)
				authRoutes.PUT("/:id", localHandler.UpdateLocal)
				authRoutes.DELETE("/:id", localHandler.DeleteLocal)
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
