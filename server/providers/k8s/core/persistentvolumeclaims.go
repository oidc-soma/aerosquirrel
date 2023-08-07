package core

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers/k8s"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FetchPersistentVolumeClaims(ctx context.Context, p k8s.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	var config v1.ListOptions

	for {
		res, err := p.Client.CoreV1().PersistentVolumeClaims("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, pvc := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range pvc.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, &models.Resource{
				TeamId:   teamId,
				Name:     pvc.Name,
				Type:     "PersistentVolumeClaim",
				Cost:     0,
				Region:   pvc.Namespace,
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