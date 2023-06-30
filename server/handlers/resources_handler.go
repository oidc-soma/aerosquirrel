package handlers

import (
	"context"
	"encoding/json"
	database "github.com/oidc-soma/aerosquirrel/server/database/mongo"
	"github.com/oidc-soma/aerosquirrel/server/models"
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

func (h *ApiHandler) CreateResource(c *gin.Context) {
	var resource models.Resource

	err := json.NewDecoder(c.Request.Body).Decode(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.CreateResource(context.Background(), &resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resource)
}

func (h *ApiHandler) GetAllResources(c *gin.Context) {

	resources, err := h.db.FindAllResources(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}

func (h *ApiHandler) GetOneResource(c *gin.Context) {
	resource, err := h.db.FindOneResource(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *ApiHandler) DeleteOneResource(c *gin.Context) {
	err := h.db.DeleteOneResource(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
