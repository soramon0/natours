package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/soramon0/natrous/pkg/database"
	"github.com/soramon0/natrous/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByID(id string) (*models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	var user *models.User
	collection := database.GetCollection("users")

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

func GetUsers() (*[]models.User, error) {
	var users []models.User
	collection := database.GetCollection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to retrieve users")
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to retrieve users")
		}

		users = append(users, singleUser)
	}

	return &users, nil
}

func CreateUser() (*models.User, error) {
	collection := database.GetCollection("users")
	newUser := models.User{
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
