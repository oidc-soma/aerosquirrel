package k8s

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *Provider) FetchPodResources(ctx context.Context, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	var config v1.ListOptions

	for {
		res, err := p.client.CoreV1().Pods("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, pod := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range pod.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, &models.Resource{
				TeamId: teamId,
				Name:   pod.Name,
				Type:   "pod",
				Cost:   0,
				Metadata: map[string]string{
					"namespace": pod.Namespace,
				},
				Tags:     tags,
				Link:     "",
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
