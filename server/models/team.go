// Package models provides models for the application.
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Team is a team.
type Team struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	OwnerId primitive.ObjectID `json:"owner" bson:"owner"`
}
