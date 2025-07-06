package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"local-gems-server/internal/app/usecases"
	"local-gems-server/internal/core/errors"
)

type LocalHandler struct {
	localUsecase usecases.LocalUsecase
}

func NewLocalHandler(usecase usecases.LocalUsecase) *LocalHandler {
	return &LocalHandler{
		localUsecase: usecase,
	}
}

func (h *LocalHandler) SearchLocals(c *gin.Context) {
	query := c.Query("q")

	locals, err := h.localUsecase.SearchLocals(query)
	if err != nil {
		if _, ok := err.(*errors.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locals)
}
