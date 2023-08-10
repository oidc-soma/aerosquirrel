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

// Package oracledatabase provides a client for interacting with the Oracle Cloud Infrastructure database service.
package oracledatabase

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchAutonomousDatabases fetches autonomous databases from the Oracle Cloud Infrastructure.
func FetchAutonomousDatabases(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	databaseClient, err := database.NewDatabaseClientWithConfigurationProvider(p.Client)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.Client.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := database.ListAutonomousDatabasesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := databaseClient.ListAutonomousDatabases(ctx, config)
	if err != nil {
		return resources, err
	}

	for _, autonomousDatabase := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range autonomousDatabase.FreeformTags {
			tags = append(tags, models.Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err := p.Client.Region()
		if err != nil {
			return resources, err
		}

		resources = append(resources, &models.Resource{
			TeamId:   teamId,
			Name:     *autonomousDatabase.DisplayName,
			Type:     "Autonomous Database",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   region,
		})
	}

	return resources, nil
}
