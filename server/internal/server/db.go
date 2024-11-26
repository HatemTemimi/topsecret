package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"server/config"
	"server/internal/rental/types"

	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func randomLatLngInTunis() (string, string) {
	rand.Seed(time.Now().UnixNano())
	lat := 36.74 + rand.Float64()*(36.88-36.74) // Latitude: 36.74 to 36.88
	lng := 10.10 + rand.Float64()*(10.29-10.10) // Longitude: 10.10 to 10.29
	return fmt.Sprintf("%.5f", lat), fmt.Sprintf("%.5f", lng)
}

func (db *DB) InitMockRentals() error {
	// Seed the faker to ensure random data
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	var rentals []types.Rental
	for i := 0; i < 20; i++ {
		var lat, lng = randomLatLngInTunis()
		rental := types.Rental{
			Name:         gofakeit.Company(),
			StreetNumber: gofakeit.StreetNumber(),
			Street:       gofakeit.StreetName(),
			City:         gofakeit.City(),
			Country:      gofakeit.Country(),
			FullAddress:  gofakeit.Address().Address,
			Lat:          lat,
			Lng:          lng,
			Price:        int64(gofakeit.Number(500, 2000)),
			Bedrooms:     int64(gofakeit.Number(1, 5)),
			Bathrooms:    int64(gofakeit.Number(1, 3)),
			AreaSize:     int64(gofakeit.Number(50, 150)),
			Available:    gofakeit.Bool(),
			Tags:         strings.Split(gofakeit.Word(), ""),
			Description:  gofakeit.Sentence(10),
			Images:       []string{"https://cdn.vuetifyjs.com/images/cards/hotel.jpg"},
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		rentals = append(rentals, rental)
	}

	collection := db.GetCollection("rentals")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert rentals to interface slice
	var docs []interface{}
	for _, rental := range rentals {
		docs = append(docs, rental)
	}

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		log.Printf("Failed to insert mock rentals: %v", err)
		return err
	}

	log.Println("Mock rentals added successfully.")
	return nil
}

func (db *DB) InitAdminUser() error {
	collection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the admin user already exists
	email := "hatem.altemimi@gmail.com"
	filter := bson.M{"email": email}
	var existingUser bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		log.Println("Admin user already exists.")
		return nil
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Failed to query admin user: %v", err)
		return err
	}

	// Hash the password
	password := "11111111"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	// Create the admin user document
	adminUser := bson.M{
		"_id":       primitive.NewObjectID(),
		"firstName": "Hatem",
		"lastName":  "Temimi",
		"email":     email,
		"password":  string(hashedPassword),
		"role":      "admin",
		"phone":     "",
		"address":   "",
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
		"active":    true,
	}

	// Insert the admin user into the collection
	_, err = collection.InsertOne(ctx, adminUser)
	if err != nil {
		log.Printf("Failed to insert admin user: %v", err)
		return err
	}

	log.Println("Admin user created successfully.")
	return nil
}

// Close disconnects the MongoDB client.
func (db *DB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}
