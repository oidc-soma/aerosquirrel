/*
 * Copyright 2023 The AeroSquirrel Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package core provides a client for fetching resource data from Kubernetes.
package core

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FetchIngress fetches ingress from Kubernetes.
func FetchIngress(ctx context.Context, p providers.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	resources := make([]*models.Resource, 0)

	var config v1.ListOptions

	for {
		res, err := p.K8sClient.NetworkingV1().Ingresses("").List(ctx, config)
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
