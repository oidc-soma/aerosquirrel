package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"net/http"
)

func (h *ApiHandler) CreateUser(c *gin.Context) {
	var user models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
