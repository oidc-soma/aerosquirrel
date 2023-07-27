package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	database "github.com/oidc-soma/aerosquirrel/server/database/mongo"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"reflect"
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
	var resources []*models.Resource
	var err error

	resources, err = h.db.FindAllResourcesByTeamId(context.Background(), c.Query("teamId"))
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

func (h *ApiHandler) GetResourcesByFilter(c *gin.Context) {
	var filters []models.Filter
	var resources []*models.Resource
	caser := cases.Title(language.AmericanEnglish)

	err := json.NewDecoder(c.Request.Body).Decode(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource := models.Resource{}
	for index, filter := range filters {
		_, ok := reflect.TypeOf(resource).FieldByName(caser.String(filter.Field))
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid field"})
			return
		}
		switch filter.Operator {
		case "eq":
			filters[index].Operator = "$eq"
		case "ne":
			filters[index].Operator = "$ne"
		case "contain":
			if filter.Field == "tags" {
				filters[index].Operator = "$in"
			} else {
				filters[index].Operator = "$regex"
			}
		case "not_contain":
			if filter.Field == "tags" {
				filters[index].Operator = "$nin"
			} else {
				filters[index].Operator = "$not"
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid operator"})
		}
	}
	resources, err = h.db.FindMultipleResources(context.Background(), user.TeamId, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
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

	existingResource, err := h.db.FindOneResource(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resource.TeamId.IsZero() {
		resource.TeamId = existingResource.TeamId
	}
	if resource.Name == "" {
		resource.Name = existingResource.Name
	}
	if resource.Type == "" {
		resource.Type = existingResource.Type
	}
	if resource.Cost == 0 {
		resource.Cost = existingResource.Cost
	}
	if resource.Metadata == nil {
		resource.Metadata = existingResource.Metadata
	}
	if resource.Tags == nil {
		resource.Tags = existingResource.Tags
	}
	if resource.Link == "" {
		resource.Link = existingResource.Link
	}
	if resource.Provider == "" {
		resource.Provider = existingResource.Provider
	}
	if resource.Region == "" {
		resource.Region = existingResource.Region
	}
	if resource.Account == "" {
		resource.Account = existingResource.Account
	}
	if resource.AccountId == "" {
		resource.AccountId = existingResource.AccountId
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

func (h *ApiHandler) ImportCSPResources(c *gin.Context) {
	csp := c.Query("csp")

	team, err := h.GetCurrentTeam(c)
	if err != nil {
		return
	}

	var resources []*models.Resource

	// TODO(krapie): CSP credentials들을 어떤 방식으로 받아 올 것인지 고민해야 함
	switch csp {
	case "oci":
		ociClient := oci.NewProvider()
		resources, err = ociClient.FetchResources(team.Id)
	case "k8s":
		k8sClient := k8s.NewProvider()
		resources, err = k8sClient.FetchResources(team.Id)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.BulkCreateResources(context.Background(), resources)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}
