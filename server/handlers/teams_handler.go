package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *ApiHandler) CreateTeam(c *gin.Context) {
	var team models.Team

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}
	team.OwnerId = userId.(primitive.ObjectID)

	err := json.NewDecoder(c.Request.Body).Decode(&team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.CreateTeam(context.Background(), &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.AddUserToTeam(context.Background(), team.Id.Hex(), userId.(primitive.ObjectID).Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func (h *ApiHandler) GetTeams(c *gin.Context) {
	teams, err := h.db.FindTeams(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *ApiHandler) GetTeam(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	team, err := h.db.FindTeam(context.Background(), c.Param("id"))
	if err != nil && err.Error() != "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.db.FindUser(context.Background(), userId.(primitive.ObjectID).Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	team, err = h.db.FindTeam(context.Background(), user.TeamId.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *ApiHandler) UpdateTeam(c *gin.Context) {
	var team models.Team
	var objectId *primitive.ObjectID
	id := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingTeam, err := h.db.FindTeam(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if team.Name == "" {
		team.Name = existingTeam.Name
	}
	if team.OwnerId.IsZero() {
		team.OwnerId = existingTeam.OwnerId
	}

	objectId, err = h.db.UpdateTeam(context.Background(), id, &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	team.Id = *objectId
	c.JSON(http.StatusOK, team)
}

func (h *ApiHandler) DeleteTeam(c *gin.Context) {
	err := h.db.DeleteTeam(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *ApiHandler) JoinTeam(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}
	teamId := c.Param("id")

	_, err := h.db.FindTeam(context.Background(), teamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.AddUserToTeam(context.Background(), teamId, userId.(primitive.ObjectID).Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
