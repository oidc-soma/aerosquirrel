package oracledatabase

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci"
	"github.com/oracle/oci-go-sdk/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchAutonomousDatabase(ctx context.Context, p oci.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
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
