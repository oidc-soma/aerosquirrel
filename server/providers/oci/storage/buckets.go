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

// Package storage contains the logic for fetching storage resources from OCI.
package storage

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"github.com/oracle/oci-go-sdk/objectstorage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchBuckets fetches buckets from OCI.
func FetchBuckets(ctx context.Context, p providers.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	objectStorageClient, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(p.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	getNamespaceRequestConfig := objectstorage.GetNamespaceRequest{
		CompartmentId: &tenancyOCID,
	}

	getNamespaceOutput, err := objectStorageClient.GetNamespace(ctx, getNamespaceRequestConfig)
	if err != nil {
		return resources, err
	}

	listBucketsRequestConfig := objectstorage.ListBucketsRequest{
		CompartmentId: &tenancyOCID,
		NamespaceName: getNamespaceOutput.Value,
	}

	output, err := objectStorageClient.ListBuckets(context.Background(), listBucketsRequestConfig)
	if err != nil {
		return resources, err
	}

	for _, bucket := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range bucket.FreeformTags {
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
			Name:     *bucket.Name,
			Type:     "ObjectStorage Bucket",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   region,
		})
	}

	return resources, nil
}
