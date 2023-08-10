package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	Client *aws.Config
}

type (
	FetchDataFunction func(ctx context.Context, p Provider, teamId primitive.ObjectID) ([]*models.Resource, error)
)
