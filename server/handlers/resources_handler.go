package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	ctx context.Context
}

func NewApiHandler(ctx context.Context) *ApiHandler {
	return &ApiHandler{
		ctx: ctx,
	}
}

func (h *ApiHandler) GetResources(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resources": "test",
	})
}
