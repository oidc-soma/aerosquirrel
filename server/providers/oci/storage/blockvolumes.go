// Package storage contains the logic for fetching storage resources from OCI.
package storage

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchBlockVolumes fetches buckets from OCI.
func FetchBlockVolumes(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	blockStorageClient, err := core.NewBlockstorageClientWithConfigurationProvider(p.Client)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.Client.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := core.ListVolumesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := blockStorageClient.ListVolumes(ctx, config)
	if err != nil {
		return resources, err
	}

	for _, volume := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range volume.FreeformTags {
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
			Name:     *volume.DisplayName,
			Type:     "Block Volume",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   region,
		})
	}

	return resources, nil
}
