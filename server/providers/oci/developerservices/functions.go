package developerservices

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/functions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchFunctions(ctx context.Context, applicationId *string, p oci.Provider, teamId primitive.ObjectID, functionsManagementClient functions.FunctionsManagementClient) ([]*models.Resource, error) {
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

		region, err := p.Client.Region()
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
