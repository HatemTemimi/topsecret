package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"server/config"
	"server/internal/rental/types"

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

func (db *DB) InitMockRentals() error {
	rentals := []types.Rental{
		{
			Name: "Rental 1", StreetNumber: "10", Street: "Avenue Habib Bourguiba", City: "Tunis",
			Country: "Tunisia", FullAddress: "10 Avenue Habib Bourguiba, Tunis, Tunisia",
			Lat: "36.81897", Lng: "10.16579", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 2", StreetNumber: "5", Street: "Rue Docteur Calmette", City: "Tunis",
			Country: "Tunisia", FullAddress: "5 Rue Dr Calmette, Tunis, Tunisia",
			Lat: "36.83196", Lng: "10.17555", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 3", StreetNumber: "15", Street: "Rue de Marseille", City: "Tunis",
			Country: "Tunisia", FullAddress: "15 Rue de Marseille, Tunis, Tunisia",
			Lat: "36.80649", Lng: "10.18153", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 4", StreetNumber: "45", Street: "Avenue de la Liberté", City: "Tunis",
			Country: "Tunisia", FullAddress: "45 Avenue de la Liberté, Tunis, Tunisia",
			Lat: "36.81492", Lng: "10.17737", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 5", StreetNumber: "8", Street: "Rue de Lyon", City: "Tunis",
			Country: "Tunisia", FullAddress: "8 Rue de Lyon, Tunis, Tunisia",
			Lat: "36.80212", Lng: "10.17685", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 6", StreetNumber: "30", Street: "Avenue Jugurtha", City: "Ariana",
			Country: "Tunisia", FullAddress: "30 Avenue Jugurtha, Ariana, Tunisia",
			Lat: "36.86681", Lng: "10.16448", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 7", StreetNumber: "12", Street: "Rue des Orangers", City: "Ariana",
			Country: "Tunisia", FullAddress: "12 Rue des Orangers, Ariana, Tunisia",
			Lat: "36.86067", Lng: "10.17039", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 8", StreetNumber: "20", Street: "Rue de l'Indépendance", City: "Ariana",
			Country: "Tunisia", FullAddress: "20 Rue de l'Indépendance, Ariana, Tunisia",
			Lat: "36.86019", Lng: "10.16489", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 9", StreetNumber: "2", Street: "Rue Mokhtar Attia", City: "Tunis",
			Country: "Tunisia", FullAddress: "2 Rue Mokhtar Attia, Tunis, Tunisia",
			Lat: "36.81733", Lng: "10.16642", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 10", StreetNumber: "34", Street: "Rue Hedi Nouira", City: "Ariana",
			Country: "Tunisia", FullAddress: "34 Rue Hedi Nouira, Ariana, Tunisia",
			Lat: "36.86431", Lng: "10.16924", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 11", StreetNumber: "17", Street: "Avenue Taieb Mhiri", City: "Tunis",
			Country: "Tunisia", FullAddress: "17 Avenue Taieb Mhiri, Tunis, Tunisia",
			Lat: "36.80761", Lng: "10.17426", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 12", StreetNumber: "22", Street: "Rue Ibn Khaldoun", City: "Ariana",
			Country: "Tunisia", FullAddress: "22 Rue Ibn Khaldoun, Ariana, Tunisia",
			Lat: "36.85791", Lng: "10.16473", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 13", StreetNumber: "5", Street: "Rue Ahmed Tlili", City: "Tunis",
			Country: "Tunisia", FullAddress: "5 Rue Ahmed Tlili, Tunis, Tunisia",
			Lat: "36.80873", Lng: "10.17244", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 14", StreetNumber: "28", Street: "Rue de Rome", City: "Tunis",
			Country: "Tunisia", FullAddress: "28 Rue de Rome, Tunis, Tunisia",
			Lat: "36.80543", Lng: "10.17652", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 15", StreetNumber: "1", Street: "Rue Mongi Slim", City: "Ariana",
			Country: "Tunisia", FullAddress: "1 Rue Mongi Slim, Ariana, Tunisia",
			Lat: "36.86047", Lng: "10.17189", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 16", StreetNumber: "9", Street: "Rue Abdelkader", City: "Tunis",
			Country: "Tunisia", FullAddress: "9 Rue Abdelkader, Tunis, Tunisia",
			Lat: "36.81267", Lng: "10.16735", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 17", StreetNumber: "3", Street: "Rue de la République", City: "Ariana",
			Country: "Tunisia", FullAddress: "3 Rue de la République, Ariana, Tunisia",
			Lat: "36.85711", Lng: "10.16433", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 18", StreetNumber: "11", Street: "Rue Saad Zaghloul", City: "Tunis",
			Country: "Tunisia", FullAddress: "11 Rue Saad Zaghloul, Tunis, Tunisia",
			Lat: "36.80617", Lng: "10.16719", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 19", StreetNumber: "6", Street: "Avenue Farhat Hached", City: "Ariana",
			Country: "Tunisia", FullAddress: "6 Avenue Farhat Hached, Ariana, Tunisia",
			Lat: "36.86321", Lng: "10.17093", Images: []string{}, Agree: false, Status: false,
		},
		{
			Name: "Rental 20", StreetNumber: "19", Street: "Rue Salah Ben Youssef", City: "Tunis",
			Country: "Tunisia", FullAddress: "19 Rue Salah Ben Youssef, Tunis, Tunisia",
			Lat: "36.81217", Lng: "10.16559", Images: []string{}, Agree: false, Status: false,
		},
	}

	//set mockup Image
	for i := range rentals {
		rentals[i].Images = []string{"https://cdn.vuetifyjs.com/images/cards/hotel.jpg"}
	}

	collection := db.GetCollection("rentals")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert rentals to an array of `interface{}` for bulk insertion
	var docs []interface{}
	for _, rental := range rentals {
		docs = append(docs, rental) // Do not include `_id` in rentals
	}

	// Insert rentals into the collection
	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		log.Printf("Failed to insert mock rentals: %v", err)
		return err
	}

	log.Println("Mock rentals added successfully to the rentals collection.")
	return nil
}

// Close disconnects the MongoDB client.
func (db *DB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}
