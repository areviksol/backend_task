package database

import (
	"context"
	"fmt"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Database struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	database := client.Database("task")
	collection := database.Collection("records")

	return &Database{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (d *Database) CheckRecord(identifier string) (bool, error) {
	id, err := strconv.ParseInt(identifier, 10, 64)
	if err != nil {
		fmt.Println("Wrong id type")
		return false, fmt.Errorf("failed to check record: %v", err) 
	}
	filter := bson.M{"identifier": id}
	count, err := d.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, fmt.Errorf("failed to check record: %v", err)
	}
	return count > 0, nil
}

func (d *Database) AddRecord(identifier string) error {
	_, err := d.collection.InsertOne(context.Background(), bson.M{"identifier": identifier})
	if err != nil {
		return fmt.Errorf("failed to add record: %v", err)
	}
	return nil
}
