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
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	database "github.com/oidc-soma/aerosquirrel/server/database/mongo"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/aws"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/pelletier/go-toml/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"reflect"
)

// ApiHandler is a handler for the API.
type ApiHandler struct {
	ctx context.Context
	db  *database.Client
}

// NewApiHandler creates a new ApiHandler.
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

// CreateResource creates a resource.
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

// GetResources gets all resources.
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

// GetOneResource gets one resource.
func (h *ApiHandler) GetOneResource(c *gin.Context) {
	resource, err := h.db.FindOneResource(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resource)
}

// UpdateResource updates a resource.
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

// UpdateResource updates a resource.
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

// DeleteOneResource deletes a resource.
func (h *ApiHandler) DeleteOneResource(c *gin.Context) {
	err := h.db.DeleteOneResource(context.Background(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// ImportCSPResources imports resources from CSP.
func (h *ApiHandler) ImportCSPResources(c *gin.Context) {
	var config models.Config
	csp := c.Query("csp")

	team, err := h.GetCurrentTeam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resources []*models.Resource

	err = toml.Unmarshal([]byte(csp), &config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(config.AWS) > 0 {
		awsProviders := make([]aws.Provider, 0)
		for _, account := range config.AWS {
			if account.Source == "CREDENTIALS_FILE" {
				if len(account.Path) > 0 {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile), awsConfig.WithSharedCredentialsFiles(
						[]string{account.Path},
					))
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
					awsProviders = append(awsProviders, aws.Provider{
						Client: &cfg,
					})
				} else {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile))
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
					awsProviders = append(awsProviders, aws.Provider{
						Client: &cfg,
					})
				}
			} else if account.Source == "ENVIRONMENT_VARIABLES" {
				cfg, err := awsConfig.LoadDefaultConfig(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				awsProviders = append(awsProviders, aws.Provider{
					Client: &cfg,
				})
			}
		}
		resources, err = aws.FetchResources(awsProviders, team.Id)
	} else if len(config.OCI) > 0 {
		ociProviders := make([]oci.Provider, 0)
		for _, account := range config.OCI {
			if account.Source == "CREDENTIALS_FILE" {
				client := common.DefaultConfigProvider()
				ociProviders = append(ociProviders, oci.Provider{
					Client: client,
				})
			}
		}
		resources, err = oci.FetchResources(ociProviders, team.Id)
	} else if len(config.Kubernetes) > 0 {
		k8sProviders := make([]k8s.Provider, 0)
		for _, account := range config.Kubernetes {
			kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
				&clientcmd.ClientConfigLoadingRules{ExplicitPath: account.Path},
				&clientcmd.ConfigOverrides{}).ClientConfig()
			if err != nil {
				log.Fatal(err)
			}

			client, err := kubernetes.NewForConfig(kubeConfig)
			if err != nil {
				log.Fatal(err)
			}

			k8sProviders = append(k8sProviders, k8s.Provider{
				Client: client,
			})
		}
		resources, err = k8s.FetchResources(k8sProviders, team.Id)
	}

	err = h.db.BulkCreateResources(context.Background(), resources)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}
