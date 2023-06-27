package handlers

import (
	"context"
	database "github.com/oidc-soma/aerosquirrel/server/database/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	ctx context.Context
	db  *database.Client
}

func NewApiHandler(ctx context.Context) *ApiHandler {
	db, err := database.Dial(ctx)
	if err != nil {
		panic(err)
	}

	return &ApiHandler{
		ctx: ctx,
		db:  db,
	}
}

func (h *ApiHandler) GetResources(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resources": "test",
	})
}
