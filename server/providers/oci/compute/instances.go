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

// Package compute provides a client for interacting with the Oracle Cloud Infrastructure Compute service.
package compute

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"github.com/oracle/oci-go-sdk/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchInstances fetches instances from the Oracle Cloud Infrastructure Compute service.
func FetchInstances(ctx context.Context, p providers.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	computeClient, err := core.NewComputeClientWithConfigurationProvider(p.OciClient)
	if err != nil {
		return nil, err
	}

	tenancyOCID, err := p.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := core.ListInstancesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := computeClient.ListInstances(ctx, config)
	if err != nil {
		return resources, err
	}

	for _, instance := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range instance.FreeformTags {
			tags = append(tags, models.Tag{
				Key:   key,
				Value: value,
			})
		}

		resources = append(resources, &models.Resource{
			TeamId:   teamId,
			Name:     *instance.DisplayName,
			Type:     "Instance",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   *instance.Region,
		})
	}

	return resources, nil
}
