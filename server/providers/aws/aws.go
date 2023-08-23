package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/aws/ec2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	Client *aws.Config
}

type (
	FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)
)

var supportedServices = []FetchDataFunction{
	ec2.FetchInstances,
}

func FetchResources(providers []Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()

	for _, provider := range providers {
		for _, service := range supportedServices {
			awsResources, err := service(ctx, provider, teamId)
			if err != nil {
				return resources, err
			}
			resources = append(resources, awsResources...)
		}
	}

	return resources, nil
}
