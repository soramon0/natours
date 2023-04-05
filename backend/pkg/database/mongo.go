package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnection(uri string, l *log.Logger) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	// Check that MongoDB server has been found and connected to
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	return client
}

func CloseConnection(c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.Disconnect(ctx)
}

func GetCollection(c *mongo.Client, collectionName string) *mongo.Collection {
	v := os.Getenv("DB_NAME")
	if v == "" {
		panic("DB_NAME env variable was not defined")
	}
	return c.Database(v).Collection(collectionName)
}
