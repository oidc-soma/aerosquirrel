package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resource struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamId    primitive.ObjectID `json:"teamId" bson:"teamId"`
	Name      string             `json:"name" bson:"name"`
	Type      string             `json:"type" bson:"type"`
	Cost      float64            `json:"cost" bson:"cost"`
	Metadata  map[string]string  `json:"metadata" bson:"metadata"`
	Tags      []Tag              `json:"tags" bson:"tags"`
	Link      string             `json:"link" bson:"link"`
	Provider  string             `json:"provider" bson:"provider"`
	Region    string             `json:"region" bson:"region"`
	Account   string             `json:"account" bson:"account"`
	AccountId string             `json:"accountId" bson:"accountId"`
}

type Tag struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
