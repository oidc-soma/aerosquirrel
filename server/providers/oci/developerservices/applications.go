package developerservices

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/functions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchApplications(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	functionsManagementClient, err := functions.NewFunctionsManagementClientWithConfigurationProvider(p.Client)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := p.Client.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := functions.ListApplicationsRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := functionsManagementClient.ListApplications(ctx, config)
	if err != nil {
		return resources, err
	}

	for _, application := range output.Items {
		tags := make([]models.Tag, 0)

		for key, value := range application.FreeformTags {
			tags = append(tags, models.Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err := p.Client.Region()
		if err != nil {
			return resources, err
		}

		allFunctions, err := FetchFunctions(ctx, application.Id, p, teamId, functionsManagementClient)
		if err != nil {
			return resources, err
		}
		if len(allFunctions) > 0 {
			resources = append(resources, allFunctions...)
		}

		resources = append(resources, &models.Resource{
			TeamId:   teamId,
			Name:     *application.DisplayName,
			Type:     "Application",
			Cost:     0,
			Tags:     tags,
			Link:     "",
			Provider: "OCI",
			Region:   region,
		})
	}

	return resources, nil
}
