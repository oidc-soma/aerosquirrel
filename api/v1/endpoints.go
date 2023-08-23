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

// Package v1 provides endpoints for the AeroSquirrel API.
package v1

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oidc-soma/aerosquirrel/server/handlers"
)

// Endpoints returns a router with all endpoints.
func Endpoints(ctx context.Context) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST, PUT, PATCH, DELETE, OPTIONS"},
		AllowHeaders:     []string{"Origin, Content-Type, Content-Length, Accept-Encoding, Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

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
		v1.POST("/resources/search/", apiHandler.GetResourcesByFilter)
		v1.PATCH("/resources/:id", apiHandler.UpdateOneResource)
		v1.DELETE("/resources/:id", apiHandler.DeleteOneResource)

		v1.GET("/resources/import", apiHandler.ImportCSPResources)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	return router
}
