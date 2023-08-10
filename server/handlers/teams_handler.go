/*
 * Copyright 2023 The AeroSquirrel Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package handlers provides handlers for the application.
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateTeam creates a team.
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

// GetTeams gets all teams.
func (h *ApiHandler) GetTeams(c *gin.Context) {
	teams, err := h.db.FindTeams(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

// GetTeam gets a team.
func (h *ApiHandler) GetTeam(c *gin.Context) {
	team, err := h.db.FindTeam(context.Background(), c.Param("id"))
	if err != nil && err.Error() != "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	team, err = h.GetCurrentTeam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

// GetCurrentTeam gets the current team from the context.
func (h *ApiHandler) GetCurrentTeam(c *gin.Context) (*models.Team, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return nil, fmt.Errorf("user id not found")
	}

	user, err := h.db.FindUser(context.Background(), userId.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}

	team, err := h.db.FindTeam(context.Background(), user.TeamId.Hex())
	if err != nil {
		return nil, err
	}

	return team, nil
}

// UpdateTeam updates a team.
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

// DeleteTeam deletes a team.
func (h *ApiHandler) DeleteTeam(c *gin.Context) {
	err := h.db.DeleteTeam(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// JoinTeam joins a team.
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
