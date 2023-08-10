// Package k8s provides a client for interacting with Kubernetes.
package k8s

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"k8s.io/client-go/kubernetes"
)

// Provider is a client for interacting with Kubernetes.
type Provider struct {
	Client *kubernetes.Clientset
}

// FetchDataFunction is a function that fetches data from Kubernetes.
type FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)

var supportedServices = []FetchDataFunction{
	core.FetchDeployments,
	core.FetchIngress,
	core.FetchPersistentVolumeClaims,
	core.FetchPersistentVolumes,
	core.FetchPods,
	core.FetchServiceAccounts,
	core.FetchServices,
}

// NewProvider creates a new Provider.
func NewProvider() *Provider {
	return &Provider{
		Client: nil,
	}
}

// FetchResources fetches resources from Kubernetes.
func (p *Provider) FetchResources(teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)
	ctx := context.Background()

	for _, service := range supportedServices {
		k8sResources, err := service(ctx, *p, teamId)
		if err != nil {
			return resources, err
		}
		resources = append(resources, k8sResources...)
	}

	return resources, nil
}
