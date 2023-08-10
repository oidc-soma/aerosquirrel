// Package storage contains the logic for fetching storage resources from OCI.
package storage

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/objectstorage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchBuckets fetches buckets from OCI.
func FetchBuckets(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	objectStorageClient, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(p.Client)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.Client.TenancyOCID()
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

		region, err := p.Client.Region()
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
