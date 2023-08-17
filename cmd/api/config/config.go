// config/config.go

package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config represents the application configuration
type Config struct {
	MongoURI string
}

// LoadEnv loads environment variables from a .env file
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// GetMongoDBConnection returns a MongoDB client connection
func GetMongoDBConnection() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	return client, nil
}

// GetConfig loads the application configuration from environment variables
func GetConfig() (*Config, error) {
	if err := LoadEnv(); err != nil {
		return nil, err
	}

	mongoURI := os.Getenv("MONGO_URI")

	config := &Config{
		MongoURI: mongoURI,
	}

	return config, nil
}
