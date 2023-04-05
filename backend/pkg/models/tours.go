package models

import (
	"context"
	"fmt"
	"natours/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tour struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	Duration        int                `bson:"duration,omitempty" json:"duration,omitempty"`
	MaxGroupSize    int                `bson:"maxGroupSize,omitempty" json:"maxGroupSize,omitempty"`
	Difficulty      string             `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	RatingsAverage  float64            `bson:"ratingsAverage,omitempty" json:"ratingsAverage,omitempty"`
	RatingsQuantity int                `bson:"ratingsQuantity,omitempty" json:"ratingsQuantity,omitempty"`
	Price           int                `bson:"price,omitempty" json:"price,omitempty"`
	Summary         string             `bson:"summary,omitempty" json:"summary,omitempty"`
	Description     string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageCover      string             `bson:"imageCover,omitempty" json:"imageCover,omitempty"`
	Images          []string           `bson:"images,omitempty" json:"images,omitempty"`
	StartDates      []string           `bson:"startDates,omitempty" json:"startDates,omitempty"`
}

type TourService interface {
	// Methods for querying tours
	ByID(id string) (*Tour, error)
	Find() ([]*Tour, error)
	// ByEmail(email string) (*Tour, error)

	// Methods for altering tours
	Create(payload CreateTourPayload) (*Tour, error)
	// Update(tour *Tour) error
	// Delete(id string) error
}

type tourService struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewTourService(client *mongo.Client) TourService {
	return &tourService{
		client: client,
		coll:   database.GetCollection(client, "tours"),
	}
}

func (ts *tourService) ByID(id string) (*Tour, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid tour id")
	}

	var tour Tour
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ts.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&tour)

	return &tour, err
}

func (ts *tourService) Find() ([]*Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := ts.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	tours := []*Tour{}
	for cursor.Next(ctx) {
		var singleTour *Tour
		if err = cursor.Decode(&singleTour); err != nil {
			return nil, err
		}

		tours = append(tours, singleTour)
	}

	return tours, nil
}

type CreateTourPayload struct {
	Name  string `json:"name" validate:"required,omitempty"`
	Price int    `json:"price" validate:"required,number,gte=1,omitempty"`
}

func (ts *tourService) Create(payload CreateTourPayload) (*Tour, error) {
	tour := &Tour{
		Id:    primitive.NewObjectID(),
		Name:  payload.Name,
		Price: payload.Price,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := ts.coll.InsertOne(ctx, tour)

	return tour, err
}
