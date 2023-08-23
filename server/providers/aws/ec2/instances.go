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

// Package ec2 provides a client for interacting with EC2.
package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"github.com/oidc-soma/aerosquirrel/server/providers"
	"github.com/oidc-soma/aerosquirrel/server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

// FetchInstances fetches EC2 instances from AWS.
func FetchInstances(ctx context.Context, p providers.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	var nextToken string
	resources := make([]*models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*p.AWSClient)

	oldRegion := p.AWSClient.Region
	p.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*p.AWSClient)
	p.AWSClient.Region = oldRegion

	for {
		output, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			NextToken: &nextToken,
		})
		if err != nil {
			return resources, err
		}

		for _, reservations := range output.Reservations {
			for _, instance := range reservations.Instances {
				tags := make([]models.Tag, 0)

				name := ""
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						name = *tag.Value
					}
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}

				monthlyCost := 0.0

				if instance.State.Name != "stopped" {
					// no need to calc usage and fetch price if stopped

					startOfMonth := utils.BeginningOfMonth(time.Now())
					hourlyUsage := 0
					if instance.LaunchTime.Before(startOfMonth) {
						hourlyUsage = int(time.Since(startOfMonth).Hours())
					} else {
						hourlyUsage = int(time.Since(*instance.LaunchTime).Hours())
					}

					pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
						ServiceCode: aws.String("AmazonEC2"),
						Filters: []types.Filter{
							{
								Field: aws.String("operatingSystem"),
								Value: aws.String("linux"),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("instanceType"),
								Value: aws.String(string(instance.InstanceType)),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("regionCode"),
								Value: aws.String(p.AWSClient.Region),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("capacitystatus"),
								Value: aws.String("Used"),
								Type:  types.FilterTypeTermMatch,
							},
						},
						MaxResults: aws.Int32(1),
					})
					if err != nil {
						return resources, err
					}

					hourlyCost := 0.0

					if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {

						priceReport := models.PriceReport{}
						err := json.Unmarshal([]byte(pricingOutput.PriceList[0]), &priceReport)
						if err != nil {
							return resources, err
						}

						for _, onDemand := range priceReport.Terms.OnDemand {
							for _, priceDimension := range onDemand.PriceDimensions {
								hourlyCost, err = strconv.ParseFloat(priceDimension.PricePerUnit.USD, 64)
								if err != nil {
									return resources, err
								}
								break
							}
							break
						}

					}

					monthlyCost = float64(hourlyUsage) * hourlyCost

				}

				resources = append(resources, &models.Resource{
					TeamId: teamId,
					Name:   name,
					Type:   "EC2",
					Cost:   monthlyCost,
					Metadata: map[string]string{
						"instanceType": string(instance.InstanceType),
						"state":        string(instance.State.Name),
					},
					Region:   p.AWSClient.Region,
					Tags:     tags,
					Link:     fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", p.AWSClient.Region, p.AWSClient.Region, *instance.InstanceId),
					Provider: "AWS",
				})
			}
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		nextToken = *output.NextToken
	}

	return resources, nil
}
