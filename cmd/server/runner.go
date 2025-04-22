package main
import (
	"github.com/gin-gonic/gin"
	"localgems/config"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.Place{}, &model.Photo{}, &model.Review{})

	r := gin.Default()

	r.GET("/places", handler.GetPlaces)
	r.GET("/places/:id", handler.GetPlaceByID)

	r.Run(":8080")
}