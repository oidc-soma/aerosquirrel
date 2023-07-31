package oci

import (
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oracle/oci-go-sdk/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	Client common.ConfigurationProvider
}

func NewProvider() *Provider {
	return &Provider{
		Client: common.DefaultConfigProvider(),
	}
}

func (p *Provider) FetchResources(teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	return resources, nil
}
