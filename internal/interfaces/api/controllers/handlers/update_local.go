package handlers

import (
	"local-gems-server/internal/core/entity"
	"local-gems-server/internal/core/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *LocalHandler) UpdateLocal(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var local entity.Local
	if err := c.ShouldBindJSON(&local); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.localUsecase.UpdateLocal(id, &local)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *errors.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	local.ID = id
	c.JSON(http.StatusOK, local)
}
