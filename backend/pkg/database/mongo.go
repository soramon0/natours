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

	return client
}

func CloseConnection(c *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Disconnect(ctx); err != nil {
		log.Fatalln(err)
	}
}

func GetCollection(c *mongo.Client, collectionName string) *mongo.Collection {
	return c.Database(os.Getenv("DB_NAME")).Collection(collectionName)
}
