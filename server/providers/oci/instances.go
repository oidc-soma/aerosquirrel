package oci

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Provider) FetchInstanceResources(ctx context.Context, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	computeClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		return nil, err
	}

	tenancyOCID, err := p.client.TenancyOCID()
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
			Type:     "OCI Instance",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   *instance.Region,
		})
	}

	return resources, nil
}
