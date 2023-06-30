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

	apiHandler := handlers.NewApiHandler(ctx)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/resources", apiHandler.CreateResource)
		v1.GET("/resources", apiHandler.GetAllResources)
		v1.GET("/resources/:id", apiHandler.GetOneResource)
		v1.DELETE("/resources/:id", apiHandler.DeleteOneResource)
		v1.POST("/users", apiHandler.CreateUser)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	return router
}
