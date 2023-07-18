package models

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamId   primitive.ObjectID `json:"teamId" bson:"teamId"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type Claims struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	jwt.RegisteredClaims
}
