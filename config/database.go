package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client instance
var mongoClient *mongo.Client

// Connect to MongoDB instance
func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in .env")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed!: %v", err)
	}
	mongoClient = client
}

func GetDB() *mongo.Database {
	dbname := os.Getenv("MONGO_DB_NAME")
	if dbname == "" {
		log.Fatal("MONGO_DB_NAME not set in .env")
	}

	return mongoClient.Database(dbname)
}
