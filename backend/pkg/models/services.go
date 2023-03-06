package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	User UserService
	Tour TourService
}

func NewServices(c *mongo.Client) *Services {
	us := NewUserService(c)
	ts := NewTourService(c)

	return &Services{
		User: us,
		Tour: ts,
	}
}
