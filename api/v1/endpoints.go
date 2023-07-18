package v1

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/handlers"
)

func Endpoints(ctx context.Context) *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	apiHandler := handlers.NewApiHandler(ctx)

	v1Public := router.Group("/api/v1")
	{
		v1Public.POST("/users", apiHandler.CreateUser)
		v1Public.POST("/login", apiHandler.Login)
	}

	v1 := router.Group("/api/v1")
	v1.Use(middleware.TokenAuth)
	{
		v1.POST("/teams", apiHandler.CreateTeam)
		v1.GET("/teams", apiHandler.GetTeams)
		v1.GET("/teams/:id", apiHandler.GetTeam)
		v1.PATCH("/teams/:id", apiHandler.UpdateTeam)
		v1.DELETE("/teams/:id", apiHandler.DeleteTeam)
		v1.GET("/teams/:id/join", apiHandler.JoinTeam)

		v1.GET("/users/:id", apiHandler.GetUser)
		v1.GET("/users", apiHandler.GetUsers)
		v1.PATCH("/users/:id", apiHandler.UpdateUser)
		v1.DELETE("/users/:id", apiHandler.DeleteUser)

		v1.POST("/resources", apiHandler.CreateResource)
		v1.GET("/resources", apiHandler.GetResources)
		v1.GET("/resources/:id", apiHandler.GetOneResource)
		v1.PATCH("/resources/:id", apiHandler.UpdateOneResource)
		v1.DELETE("/resources/:id", apiHandler.DeleteOneResource)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	return router
}
