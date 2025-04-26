package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

"localgems/internal/app/usecases"
"localgems/internal/core/errors"

)

type CoffeeHandler struct {
	coffeeUsecase usecases.CoffeeUsecase
}

func NewCoffeeHandler(usecase usecases.CoffeeUsecase) *CoffeeHandler {
	return &CoffeeHandler{
		coffeeUsecase: usecase,
	}
}

func (h *CoffeeHandler) SearchCoffees(c *gin.Context) {
	query := c.Query("q")
	
	coffees, err := h.coffeeUsecase.SearchCoffees(query)
	if err != nil {
		if _, ok := err.(*errors.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coffees)
}
