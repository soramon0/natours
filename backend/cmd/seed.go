package main

import (
	"context"
	"encoding/json"
	"natours/pkg/database"
	"natours/pkg/models"
	"natours/pkg/utils"
	"os"
	"path"
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger := utils.InitLogger()
	client := database.OpenConnection(utils.GetDatabaseBindAdress(), logger)
	defer func() {
		utils.Must(database.CloseConnection(client))
	}()

	db := client.Database("natours")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := insertUsers(db.Collection("users")); err != nil {
			logger.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		if _, err := insertTours(db.Collection("tours")); err != nil {
			logger.Fatalln(err)
		}
	}()

	wg.Wait()
}

func insertTours(coll *mongo.Collection) ([]primitive.ObjectID, error) {
	var tours []models.Tour
	var data []interface{}
	if err := readDocs("tours.json", &tours); err != nil {
		return nil, err
	}
	var ids = make([]primitive.ObjectID, 0, len(tours))
	for _, tour := range tours {
		id, _ := primitive.ObjectIDFromHex(tour.Id.String())
		tour.Id = id
		data = append(data, tour)
		ids = append(ids, id)
	}
	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func insertUsers(coll *mongo.Collection) ([]primitive.ObjectID, error) {
	var users []models.User
	var data []interface{}
	if err := readDocs("users.json", &users); err != nil {
		return nil, err
	}
	var ids = make([]primitive.ObjectID, 0, len(users))
	for _, user := range users {
		id, _ := primitive.ObjectIDFromHex(user.Id.String())
		user.Id = id
		data = append(data, user)
		ids = append(ids, id)
	}
	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func readDocs(filename string, dst any) error {
	blob, err := os.ReadFile(path.Join("data", filename))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(blob, dst); err != nil {
		return err
	}
	return nil
}
