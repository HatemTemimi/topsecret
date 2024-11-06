package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"server/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds the MongoDB client and config.
type DB struct {
	client   *mongo.Client
	database *mongo.Database
	config   config.Config
}

// NewDB initializes a new MongoDB connection.
func NewDB(cfg *config.Config) (*DB, error) {
	// Build MongoDB URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort)
	clientOptions := options.Client().ApplyURI(uri)

	// Set up a timeout context for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping to verify the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Printf("Successfully connected and pinged database: %s", cfg.DatabaseName)

	// Return a DB instance with a connected client and the specified database
	return &DB{
		client:   client,
		database: client.Database(cfg.DatabaseName),
		config:   *cfg,
	}, nil
}

// GetCollection returns a MongoDB collection.
func (db *DB) GetCollection(name string) *mongo.Collection {
	return db.database.Collection(name)
}

// Close disconnects the MongoDB client.
func (db *DB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}
