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

// Package providers provides a client for interacting with cloud providers.
package providers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oracle/oci-go-sdk/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"k8s.io/client-go/kubernetes"
)

// FetchDataFunction is a function that fetches data from the Oracle Cloud Infrastructure.
type FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)

// Provider is a struct that contains all the clients for a provider.
type Provider struct {
	AWSClient *aws.Config
	OciClient common.ConfigurationProvider
	K8sClient *kubernetes.Clientset
}
