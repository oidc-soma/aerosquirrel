package oci

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/compute"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/developerservices"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/iam"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/oracledatabase"
	"github.com/oidc-soma/aerosquirrel/server/providers/oci/storage"
	"github.com/oracle/oci-go-sdk/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	Client common.ConfigurationProvider
}

type FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)

var supportedServices = []FetchDataFunction{
	compute.FetchInstances,
	iam.FetchPolicies,
	oracledatabase.FetchAutonomousDatabases,
	storage.FetchBuckets,
	storage.FetchBlockVolumes,
	developerservices.FetchApplications,
}

func NewProvider() *Provider {
	return &Provider{
		Client: common.DefaultConfigProvider(),
	}
}

func (p *Provider) FetchResources(teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()

	for _, service := range supportedServices {
		ociResources, err := service(ctx, *p, teamId)
		if err != nil {
			return resources, err
		}
		resources = append(resources, ociResources...)
	}

	return resources, nil
}
