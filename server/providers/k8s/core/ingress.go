// Package core provides a client for fetching resource data from Kubernetes.
package core

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FetchIngress fetches ingress from Kubernetes.
func FetchIngress(ctx context.Context, p k8s.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	var config v1.ListOptions

	for {
		res, err := p.Client.NetworkingV1().Ingresses("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, ingress := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range ingress.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, &models.Resource{
				TeamId:   teamId,
				Name:     ingress.Name,
				Type:     "Ingress",
				Cost:     0,
				Region:   ingress.Namespace,
				Tags:     tags,
				Provider: "k8s",
			})
		}

		if res.GetContinue() == "" {
			break
		}

		config.Continue = res.GetContinue()
	}

	return resources, nil
}
