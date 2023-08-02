package iam

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/identity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchPolices(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(p.Client)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.Client.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := identity.ListPoliciesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := identityClient.ListPolicies(ctx, config)
	if err != nil {
		return resources, err
	}

	for _, policy := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range policy.FreeformTags {
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
			Name:     *policy.Name,
			Type:     "Identity Policy",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   region,
		})
	}

	return resources, nil
}
