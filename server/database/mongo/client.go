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

// Package database provides a client for interacting with the database.
package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/oidc-soma/aerosquirrel/server/models"
)

const (
	// AeroSquirrelDatabase is the name of the database.
	AeroSquirrelDatabase = "aerosquirrel"

	// TeamCollection is the name of the team collection.
	TeamCollection = "teams"
	// UserCollection is the name of the user collection.
	UserCollection = "users"
	// ResourceCollection is the name of the resource collection.
	ResourceCollection = "resources"
)

// Client is a client for interacting with the database.
type Client struct {
	client *mongo.Client
}

// Dial dials the database.
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

// CreateTeam creates a new team in the database.
func (m *Client) CreateTeam(ctx context.Context, team *models.Team) error {
	res, err := m.client.Database(AeroSquirrelDatabase).Collection(TeamCollection).InsertOne(ctx, team)
	if err != nil {
		return err
	}

	team.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

// FindTeams finds all teams in the database.
func (m *Client) FindTeams(ctx context.Context) ([]*models.Team, error) {
	var teams []*models.Team

	cursor, err := m.client.Database(AeroSquirrelDatabase).Collection(TeamCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// FindTeam finds a team in the database.
func (m *Client) FindTeam(ctx context.Context, id string) (*models.Team, error) {
	var team *models.Team

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = m.client.Database(AeroSquirrelDatabase).Collection(TeamCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

// UpdateTeam updates a team in the database.
func (m *Client) UpdateTeam(ctx context.Context, id string, team *models.Team) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(TeamCollection).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": team})
	if err != nil {
		return nil, err
	}

	return &objectId, nil
}

// DeleteTeam deletes a team in the database.
func (m *Client) DeleteTeam(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(TeamCollection).DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

// AddUserToTeam adds a user to a team in the database.
func (m *Client) AddUserToTeam(ctx context.Context, id string, userId string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	userIdObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).UpdateOne(ctx, bson.M{"_id": userIdObjectId}, bson.M{"$set": bson.M{"teamId": objectId}})
	if err != nil {
		return err
	}

	return nil
}

// CreateUser creates a new user in the database.
func (m *Client) CreateUser(ctx context.Context, user *models.User) error {
	res, err := m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

// FindUserByUsername finds a user in the database.
func (m *Client) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User

	err := m.client.Database("aerosquirrel").Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUsersByTeamId finds all users in the database.
func (m *Client) FindUsersByTeamId(ctx context.Context, teamId string) ([]*models.User, error) {
	var users []*models.User

	objectId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	cursor, err := m.client.Database(AeroSquirrelDatabase).Collection(UserCollection).Find(ctx, bson.M{"teamId": objectId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUser finds a user in the database.
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

// UpdateUser updates a user in the database.
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

// DeleteUser deletes a user in the database.
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

// CreateResource creates a resource in the database.
func (m *Client) CreateResource(ctx context.Context, resource *models.Resource) error {
	res, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).InsertOne(ctx, resource)
	if err != nil {
		return err
	}

	resource.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

// BulkCreateResources creates a resource in the database.
func (m *Client) BulkCreateResources(ctx context.Context, resources []*models.Resource) error {
	var documents []interface{}

	for _, resource := range resources {
		documents = append(documents, resource)
	}

	_, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	return nil
}

// FindAllResourcesByTeamId finds all resources in the database.
func (m *Client) FindAllResourcesByTeamId(ctx context.Context, teamId string) ([]*models.Resource, error) {
	var resources []*models.Resource

	objectId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	cursor, err := m.client.Database(AeroSquirrelDatabase).Collection(ResourceCollection).Find(ctx, bson.M{"teamId": objectId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// FindOneResource finds a resource in the database.
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

// FindMultipleResources finds multiple resources in the database.
func (m *Client) FindMultipleResources(ctx context.Context, teamId primitive.ObjectID, filters []models.Filter) ([]*models.Resource, error) {
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
	bsonFilter = append(bsonFilter, bson.M{"teamId": teamId})
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

// UpdateOneResource updates a resource in the database.
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

// DeleteOneResource deletes a resource in the database.
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
