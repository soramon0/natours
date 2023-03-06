package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	User UserService
}

func NewServices(c *mongo.Client) *Services {
	us := NewUserService(c)

	return &Services{
		User: us,
	}
}
