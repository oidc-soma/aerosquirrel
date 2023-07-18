package database

import (
	"context"
	"fmt"
	"github.com/oidc-soma/aerosquirrel/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	AeroSquirrelDatabase = "aerosquirrel"

	UserCollection     = "users"
	ResourceCollection = "resources"
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
	res, err := m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (m *Client) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User

	err := m.client.Database("aerosquirrel").Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Client) FindUser(ctx context.Context, id string) (*models.User, error) {
	var user *models.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Client) UpdateUser(ctx context.Context, id string, user *models.User) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return &objectId, nil
}

func (m *Client) DeleteUser(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

func (m *Client) CreateResource(ctx context.Context, resource *models.Resource) error {
	res, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).InsertOne(ctx, resource)
	if err != nil {
		return err
	}

	resource.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (m *Client) FindAllResources(ctx context.Context) ([]*models.Resource, error) {
	var resources []*models.Resource

	cursor, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (m *Client) FindOneResource(ctx context.Context, id string) (*models.Resource, error) {
	var resource *models.Resource

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&resource)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (m *Client) FindMultipleResources(ctx context.Context, filters []models.Filter) ([]*models.Resource, error) {
	var resources []*models.Resource

	var bsonFilter bson.A
	for _, f := range filters {
		if f.Operator == "$in" || f.Operator == "$nin" {
			bsonFilter = append(bsonFilter, bson.M{f.Field: bson.M{f.Operator: bson.A{f.Value}}})
		} else if f.Operator == "$not" {
			bsonFilter = append(bsonFilter, bson.M{f.Field: bson.M{f.Operator: bson.M{"$regex": f.Value}}})
		} else {
			bsonFilter = append(bsonFilter, bson.M{f.Field: bson.M{f.Operator: f.Value}})
		}
	}
	fmt.Println(bsonFilter)
	cursor, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).Find(ctx, bson.M{"$and": bsonFilter})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (m *Client) UpdateOneResource(ctx context.Context, id string, resource *models.Resource) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": resource})
	if err != nil {
		return nil, err
	}

	return &objectId, nil
}

func (m *Client) DeleteOneResource(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}
