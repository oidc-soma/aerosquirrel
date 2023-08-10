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
