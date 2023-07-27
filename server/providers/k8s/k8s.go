package k8s

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"k8s.io/client-go/kubernetes"
)

type Provider struct {
	client *kubernetes.Clientset
}

func NewProvider() *Provider {
	return &Provider{
		client: nil,
	}
}

func (p *Provider) FetchResources(teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()

	podResources, err := p.FetchPodResources(ctx, teamId)
	if err != nil {
		return nil, err
	}
	resources = append(resources, podResources...)

	return resources, nil
}
