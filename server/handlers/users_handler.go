package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oidc-soma/aerosquirrel/server/middleware"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *ApiHandler) CreateUser(c *gin.Context) {
	var user models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashedPassword

	err = h.db.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *ApiHandler) Login(c *gin.Context) {
	var user models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actualUser, err := h.db.FindUserByUsername(context.Background(), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashedPassword := actualUser.Password

	err = utils.CompareHashAndPassword(hashedPassword, user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &models.Claims{
		Id:       actualUser.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{Token: tokenString})
}

func (h *ApiHandler) GetUsers(c *gin.Context) {
	user, err := h.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users, err := h.db.FindUsersByTeamId(context.Background(), user.TeamId.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *ApiHandler) GetUser(c *gin.Context) {
	user, err := h.db.FindUser(context.Background(), c.Param("id"))
	if err != nil && err.Error() != "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err = h.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *ApiHandler) GetCurrentUser(c *gin.Context) (*models.User, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return nil, fmt.Errorf("user id not found")
	}

	user, err := h.db.FindUser(context.Background(), userId.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, fmt.Errorf("user id not found")
	}

	return user, nil
}

func (h *ApiHandler) UpdateUser(c *gin.Context) {
	var user models.User
	var objectId *primitive.ObjectID
	id := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := h.db.FindUser(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.TeamId.IsZero() {
		user.TeamId = existingUser.TeamId
	}
	if user.Username == "" {
		user.Username = existingUser.Username
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}
	if user.Password == "" {
		user.Password = existingUser.Password
	}

	objectId, err = h.db.UpdateUser(context.Background(), id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = *objectId
	c.JSON(http.StatusOK, user)
}

func (h *ApiHandler) DeleteUser(c *gin.Context) {
	err := h.db.DeleteUser(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
