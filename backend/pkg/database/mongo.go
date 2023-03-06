package database

import (
	"context"
	"time"

	"github.com/soramon0/natrous/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.GetDatabaseBindAdress()))
	if err != nil {
		panic(err)
	}

	// Check that MongoDB server has been found and connected to
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	utils.Must(client.Ping(ctx, readpref.Primary()))

	return client
}

func CloseConnection(c *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	utils.Must(c.Disconnect(ctx))
}

func GetCollection(c *mongo.Client, collectionName string) *mongo.Collection {
	return c.Database(utils.GetDatabaseName()).Collection(collectionName)
}
