package v1

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/handlers"
)

func Endpoints(ctx context.Context) *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	api := handlers.NewApiHandler(ctx)

	router.GET("/resources", api.GetResources)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return router
}
