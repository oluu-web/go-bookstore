// models/db.go

package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbName = os.Getenv("DB") // Your MongoDB database name

// ConnectToDB establishes a connection to the MongoDB database
func ConnectToDB() error {
	mongoURI := os.Getenv("MONGO_URI") // Load this from your environment variables
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	return nil
}

// GetDBCollection returns a reference to a collection in the MongoDB database
func GetDBCollection(collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func GetAllBooks() ([]Book, error) {
	collection := GetDBCollection("Books")

	filter := bson.D{}
	options := options.Find()

	cur, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var books []Book
	for cur.Next(context.Background()) {
		var book Book
		if err := cur.Decode(&book); err != nil {
			log.Panicln("Error decoding book: ", err)
			continue
		}

		books = append(books, book)

	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
