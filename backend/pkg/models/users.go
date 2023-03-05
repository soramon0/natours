package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
	Active   bool               `json:"active"`
	Photo    string             `json:"photo"`
	Password string             `json:"-"`
}
