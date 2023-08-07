package core

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FetchServices(ctx context.Context, p k8s.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	var config v1.ListOptions

	for {
		res, err := p.Client.CoreV1().Services("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, service := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range service.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, &models.Resource{
				TeamId:   teamId,
				Name:     service.Name,
				Type:     "ServiceAccount",
				Cost:     0,
				Region:   service.Namespace,
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
