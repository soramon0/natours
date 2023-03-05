package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func OpenConnection() *mongo.Client {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/", // eg. mongodb://user:password@mongo:27017/
		os.Getenv("DB_PORTOCAL"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln(err)
	}

	// Check that MongoDB server has been found and connected to
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}

	// Save client to global client
	Client = client

	return Client
}

func CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatalln(err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := Client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}
