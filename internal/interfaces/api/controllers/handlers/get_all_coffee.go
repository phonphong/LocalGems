package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CoffeeHandler) GetAllCoffees(c *gin.Context) {
	coffees, err := h.coffeeUsecase.GetAllCoffees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffees)
}
