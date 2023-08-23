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

// Package developerservices provides a client for interacting with the developer services.
package developerservices

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"github.com/oracle/oci-go-sdk/functions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchApplications fetches applications from the developer services.
func FetchFunctions(ctx context.Context, applicationId *string, p providers.Provider, teamId primitive.ObjectID, functionsManagementClient functions.FunctionsManagementClient) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	listFunctionsRequest := functions.ListFunctionsRequest{
		ApplicationId: applicationId,
	}

	output, err := functionsManagementClient.ListFunctions(ctx, listFunctionsRequest)
	if err != nil {
		return resources, err
	}

	for _, function := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range function.FreeformTags {
			tags = append(tags, models.Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err := p.OciClient.Region()
		if err != nil {
			return resources, err
		}

		resources = append(resources, &models.Resource{
			TeamId:   teamId,
			Name:     *function.DisplayName,
			Type:     "Function",
			Cost:     0,
			Region:   region,
			Tags:     tags,
			Provider: "OCI",
		})
	}

	return resources, nil
}
