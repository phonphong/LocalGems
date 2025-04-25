package handlers

import (
	"localgems/internal/core/entity"
	"localgems/internal/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CoffeeHandler) CreateCoffee(c *gin.Context) {
	var coffee entity.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.coffeeUsecase.CreateCoffee(&coffee)
	if err != nil {
		if _, ok := err.(*errors.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coffee.ID = id
	c.JSON(http.StatusCreated, coffee)
}
