package handlers

import (
	"localgems/internal/core/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CoffeeHandler) DeleteCoffee(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = h.coffeeUsecase.DeleteCoffee(id)
	if err != nil {
		if _, ok := err.(*errors.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "coffee deleted successfully"})
}
