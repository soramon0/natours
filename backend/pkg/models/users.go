package models

import (
	"context"
	"fmt"
	"time"

	"natours/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
	Active   bool               `bson:"active,omitempty" json:"active,omitempty"`
	Photo    string             `bson:"photo,omitempty" json:"photo,omitempty"`
	Password string             `bson:"password,omitempty" json:"-"`
}

type UserService interface {
	// Methods for querying users
	ByID(id string) (*User, error)
	Find() ([]*User, error)
	// ByEmail(email string) (*User, error)

	// Methods for altering users
	Create() (*User, error)
	// Update(user *User) error
	// Delete(id string) error
}

type userService struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserService(client *mongo.Client) UserService {
	return &userService{
		client: client,
		coll:   database.GetCollection(client, "users"),
	}
}

func (us *userService) ByID(id string) (*User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	var user *User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = us.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}

	return user, err
}

func (us *userService) Find() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := us.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := []*User{}

	for cursor.Next(ctx) {
		var singleUser *User
		if err = cursor.Decode(&singleUser); err != nil {
			return nil, err
		}
		users = append(users, singleUser)
	}

	return users, nil
}

func (us *userService) Create() (*User, error) {
	newUser := &User{
		Name: "Sora",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := us.coll.InsertOne(ctx, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}
