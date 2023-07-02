package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resource struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId     int64              `json:"user_id" bson:"user_id"`
	Name       string             `json:"name" bson:"name"`
	ResourceId string             `json:"resource_id" bson:"resource_id, unique"`
	Type       string             `json:"type" bson:"type"`
	Cost       float64            `json:"cost" bson:"cost"`
	Metadata   map[string]string  `json:"metadata" bson:"metadata"`
	Tags       []Tag              `json:"tags" bson:"tags"`
	Link       string             `json:"link" bson:"link"`
	Provider   string             `json:"provider" bson:"provider"`
	Region     string             `json:"region" bson:"region"`
	Account    string             `json:"account" bson:"account"`
	AccountId  string             `json:"account_id" bson:"account_id"`
}

type Tag struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
