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

// Package oci provides a client for interacting with the Oracle Cloud Infrastructure.
package oci

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/compute"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/developerservices"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/iam"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/oracledatabase"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/storage"
	"github.com/oracle/oci-go-sdk/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Provider is a client for interacting with the Oracle Cloud Infrastructure.
type Provider struct {
	Client common.ConfigurationProvider
}

// FetchDataFunction is a function that fetches data from the Oracle Cloud Infrastructure.
type FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)

var supportedServices = []FetchDataFunction{
	compute.FetchInstances,
	iam.FetchPolicies,
	oracledatabase.FetchAutonomousDatabases,
	storage.FetchBuckets,
	storage.FetchBlockVolumes,
	developerservices.FetchApplications,
}

// NewProvider creates a new Provider.
func NewProvider() *Provider {
	return &Provider{
		Client: common.DefaultConfigProvider(),
	}
}

// FetchResources fetches resources from the Oracle Cloud Infrastructure.
func (p *Provider) FetchResources(teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()

	for _, service := range supportedServices {
		ociResources, err := service(ctx, *p, teamId)
		if err != nil {
			return resources, err
		}
		resources = append(resources, ociResources...)
	}

	return resources, nil
}
