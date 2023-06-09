package handlers

import (
	"context"
	"encoding/json"
	database "github.com/oidc-soma/aerosquirrel/server/database/mongo"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *ApiHandler) GetResources(c *gin.Context) {
	var tags []models.Tag
	var resources []*models.Resource
	var err error

	query := c.Request.URL.Query()
	if len(query) == 0 {
		resources, err = h.db.FindAllResources(context.Background())
	} else {
		for key, values := range query {
			tags = append(tags, models.Tag{Key: key, Value: values[0]})
		}
		resources, err = h.db.FindMultipleResources(context.Background(), tags)
	}
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

func (h *ApiHandler) UpdateOneResource(c *gin.Context) {
	var resource models.Resource
	var objectId *primitive.ObjectID
	id := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	objectId, err = h.db.UpdateOneResource(context.Background(), id, &resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource.Id = *objectId
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
