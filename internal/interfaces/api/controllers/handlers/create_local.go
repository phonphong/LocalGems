package handlers

import (
	"local-gems-server/internal/core/entity"
	"local-gems-server/internal/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocalHandler) CreateLocal(c *gin.Context) {
	var local entity.Local
	if err := c.ShouldBindJSON(&local); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.localUsecase.CreateLocal(&local)
	if err != nil {
		if _, ok := err.(*errors.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	local.ID = id
	c.JSON(http.StatusCreated, local)
}
