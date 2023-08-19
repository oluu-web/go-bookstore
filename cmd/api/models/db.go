// models/db.go

package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//var dbName = os.Getenv("DB") // Your MongoDB database name

var dbName = "Bookstore"

// ConnectToDB establishes a connection to the MongoDB database
func ConnectToDB() error {
	errr := godotenv.Load()
	if errr != nil {
		log.Fatal("Error loading .env file: ", errr)
	}
	mongoURI := os.Getenv("MONGOURI") // Load this from your environment variables
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	log.Println("connected successfully")
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

func CreateNewBook(book Book) (string, error) {
	collection := GetDBCollection("Books")

	result, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(string); ok {
		return oid, nil
	}

	return "Created Successfully", fmt.Errorf("created successfully")
}

func GetBookByID(bookID string) (Book, error) {
	collection := GetDBCollection("Books")

	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return Book{}, fmt.Errorf("invalid objectid")
	}

	filter := bson.M{"_id": objectID}

	var book Book
	err = collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Book{}, fmt.Errorf("Book not found")
		}
		return Book{}, err
	}
	return book, nil
}

func GetGenres() ([]string, error) {
	collection := GetDBCollection("Books")

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$genre",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":   0,
				"genre": "$_id",
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var genres []string
	for cursor.Next(context.Background()) {
		var genre struct {
			Genre string `bson:"genre"`
		}
		if err := cursor.Decode(&genre); err != nil {
			log.Println("Error decoding genre: ", err)
			continue
		}
		genres = append(genres, genre.Genre)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func GetBooksByGenre(genre string) ([]Book, error) {
	collection := GetDBCollection("Books")

	filter := bson.M{
		"genre": genre,
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var books []Book
	for cur.Next(context.Background()) {
		var book Book
		if err := cur.Decode(&book); err != nil {
			log.Println("Error deoding book", err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func GetYears() ([]int, error) {
	collection := GetDBCollection("Books")

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$release_date",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":          0,
				"release_date": "$_id",
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var years []int
	for cursor.Next(context.Background()) {
		var year struct {
			Year int `bson:"release_date"`
		}
		if err := cursor.Decode(&year); err != nil {
			log.Println("Error decoding release_date: ", err)
			continue
		}
		years = append(years, year.Year)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return years, nil
}

func GetBooksByYear(release_date int) ([]Book, error) {
	collection := GetDBCollection("Books")

	filter := bson.M{
		"release_date": release_date,
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var books []Book
	for cur.Next(context.Background()) {
		var book Book
		if err := cur.Decode(&book); err != nil {
			log.Println("Error deoding book", err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func GetAuthors() ([]string, error) {
	collection := GetDBCollection("Books")

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$author",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":    0,
				"author": "$_id",
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var authors []string
	for cursor.Next(context.Background()) {
		var author struct {
			Author string `bson:"author"`
		}
		if err := cursor.Decode(&author); err != nil {
			log.Println("Error decoding author: ", err)
			continue
		}
		authors = append(authors, author.Author)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func GetBooksByAuthor(author string) ([]Book, error) {
	collection := GetDBCollection("Books")

	filter := bson.M{
		"author": author,
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var books []Book
	for cur.Next(context.Background()) {
		var book Book
		if err := cur.Decode(&book); err != nil {
			log.Println("Error deoding book", err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func UpdateBook(bookID string, updatedBook Book) error {
	collection := GetDBCollection("Books")

	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return fmt.Errorf("Invalid object ID format: ", err)
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"book_name":    updatedBook.BookName,
			"author":       updatedBook.Genre,
			"release_date": updatedBook.ReleaseDate,
			"genre":        updatedBook.Genre,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(bookID string) error {
	collection := GetDBCollection("Books")

	_, err := GetBookByID(bookID)
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return fmt.Errorf("Invalid object ID format")
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func CheckDuplicate(name string) (bool, error) {
	collection := GetDBCollection("Books")

	filter := bson.M{"book_name": name}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
