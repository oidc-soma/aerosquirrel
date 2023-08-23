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
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/compute"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/developerservices"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/iam"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/oracledatabase"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var supportedServices = []providers.FetchDataFunction{
	compute.FetchInstances,
	iam.FetchPolicies,
	oracledatabase.FetchAutonomousDatabases,
	storage.FetchBuckets,
	storage.FetchBlockVolumes,
	developerservices.FetchApplications,
}

// FetchResources fetches resources from the Oracle Cloud Infrastructure.
func FetchResources(providers []providers.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()
	for _, provider := range providers {
		for _, service := range supportedServices {
			ociResources, err := service(ctx, provider, teamId)
			if err != nil {
				return resources, err
			}
			resources = append(resources, ociResources...)
		}
	}

	return resources, nil
}
