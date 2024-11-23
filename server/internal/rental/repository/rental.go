package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"server/internal/rental/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RentalRepository interface {
	AddRental(ctx context.Context, rental types.Rental) error
	GetAllRentals(ctx context.Context) ([]types.Rental, error)
	GetRentalByID(ctx context.Context, id string) (*types.Rental, error)
	UpdateRental(ctx context.Context, id string, updatedData types.Rental) error
	DeleteRental(ctx context.Context, id string) error
}

type rentalRepository struct {
	collection *mongo.Collection
}

func NewRentalRepository(db *mongo.Database) RentalRepository {
	return &rentalRepository{
		collection: db.Collection("rentals"),
	}
}

// AddRental adds a new rental to the database
func (r *rentalRepository) AddRental(ctx context.Context, rental types.Rental) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Ensure the FullAddress field is populated
	if rental.FullAddress == "" {
		rental.FullAddress = rental.StreetNumber + " " + rental.Street + ", " + rental.City + ", " + rental.Country
	}

	_, err := r.collection.InsertOne(ctx, rental)
	if err != nil {
		log.Printf("Error inserting rental: %v", err)
		return err
	}
	return nil
}

// GetAllRentals retrieves all rentals from the database
func (r *rentalRepository) GetAllRentals(ctx context.Context) ([]types.Rental, error) {
	var rentals []types.Rental

	cursor, err := r.collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var rental types.Rental
		if err := cursor.Decode(&rental); err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rentals, nil
}

// GetRentalByID retrieves a rental by its ID
func (r *rentalRepository) GetRentalByID(ctx context.Context, id string) (*types.Rental, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, errors.New("invalid ID format")
	}

	var rental types.Rental
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&rental)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No rental found with the given ID
		}
		log.Printf("Error finding rental: %v", err)
		return nil, err
	}

	return &rental, nil
}

// UpdateRental updates an existing rental by its ID
func (r *rentalRepository) UpdateRental(ctx context.Context, id string, updatedData types.Rental) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return errors.New("invalid ID format")
	}

	// Ensure the FullAddress field is updated if street-related fields change
	if updatedData.FullAddress == "" {
		updatedData.FullAddress = updatedData.StreetNumber + " " + updatedData.Street + ", " + updatedData.City + ", " + updatedData.Country
	}

	update := bson.M{
		"$set": updatedData,
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		log.Printf("Error updating rental: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no rental found with the given ID")
	}

	return nil
}

// DeleteRental deletes a rental by its ID
func (r *rentalRepository) DeleteRental(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return errors.New("invalid ID format")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Printf("Error deleting rental: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no rental found with the given ID")
	}

	return nil
}
