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
	awsProvider "github.com/oidc-soma/aerosquirrel/server/providers/aws"
	"github.com/oidc-soma/aerosquirrel/server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

func FetchInstances(ctx context.Context, p awsProvider.Provider, teamId primitive.ObjectID) ([]*models.Resource, error) {
	var nextToken string
	resources := make([]*models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*p.Client)

	oldRegion := p.Client.Region
	p.Client.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*p.Client)
	p.Client.Region = oldRegion

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
								Value: aws.String(p.Client.Region),
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
					Region:   p.Client.Region,
					Tags:     tags,
					Link:     fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", p.Client.Region, p.Client.Region, *instance.InstanceId),
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
