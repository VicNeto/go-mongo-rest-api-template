package repository

import (
	"context"
	"fmt"
	conf "go-rest-mongodb/config"
	"go-rest-mongodb/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PlacesRepository Repository
type PlacesRepository struct{}

var config conf.Config
var collection = new(mongo.Collection)

// PlacesCollection Collection
const PlacesCollection = "Places"

func init() {
	config.Read()

	// Connect to DB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.URI))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(config.Database.DatabaseName).Collection(PlacesCollection)
}

// FindAll Get all Places
func (p *PlacesRepository) FindAll() ([]models.Place, error) {
	var places []models.Place

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.Place
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return places, err
}

// Insert Create a new Place
func (p *PlacesRepository) Insert(place models.Place) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, &place)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// Delete Delete an existing Place
func (p *PlacesRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return err
}
