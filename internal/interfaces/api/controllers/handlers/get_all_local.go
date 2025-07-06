package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocalHandler) GetAllLocals(c *gin.Context) {
	locals, err := h.localUsecase.GetAllLocals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locals)
}
