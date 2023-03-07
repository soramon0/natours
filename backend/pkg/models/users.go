package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/soramon0/natrous/pkg/database"
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
	Find() (*[]User, error)
	// ByEmail(email string) (*User, error)

	// Methods for altering users
	Create() (*User, error)
	// Update(user *User) error
	// Delete(id string) error
}

type userService struct {
	client *mongo.Client
}

func NewUserService(client *mongo.Client) UserService {
	return &userService{
		client: client,
	}
}

func (us *userService) ByID(id string) (*User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	var user *User
	collection := database.GetCollection(us.client, "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err == nil {
		return user, nil
	}

	log.Println(err)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	}

	return nil, err
}

func (us *userService) Find() (*[]User, error) {
	users := make([]User, 0)
	collection := database.GetCollection(us.client, "users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to retrieve users")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var singleUser User
		if err = cursor.Decode(&singleUser); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to retrieve users")
		}

		users = append(users, singleUser)
	}

	return &users, nil
}

func (us *userService) Create() (*User, error) {
	collection := database.GetCollection(us.client, "users")
	newUser := User{
		Id:   primitive.NewObjectID(),
		Name: "Sora",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return &newUser, nil
}
