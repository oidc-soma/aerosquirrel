package database

import (
	"context"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func Dial(ctx context.Context, uri ...string) (*Client, error) {
	URI := "mongodb://localhost:27017"
	if len(uri) != 0 {
		URI = uri[0]
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func Close(ctx context.Context, m *Client) error {
	return m.client.Disconnect(ctx)
}

func (m *Client) CreateUser(ctx context.Context, user *models.User) error {
	res, err := m.client.Database("aerosquirrel").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (m *Client) CreateResource(ctx context.Context, resource *models.Resource) error {
	res, err := m.client.Database("aerosquirrel").Collection("resources").InsertOne(ctx, resource)
	if err != nil {
		return err
	}

	resource.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (m *Client) FindAllResources(ctx context.Context) ([]*models.Resource, error) {
	var resources []*models.Resource

	cursor, err := m.client.Database("aerosquirrel").Collection("resources").Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
