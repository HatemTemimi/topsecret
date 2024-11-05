// repository/rental_repository.go
package repository

import (
	"context"
	"log"
	"time"

	"server/internal/rental/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RentalRepository interface {
	AddRental(ctx context.Context, rental types.Rental) error
	GetAllRentals(ctx context.Context) ([]types.Rental, error)
}

type rentalRepository struct {
	collection *mongo.Collection
}

func NewRentalRepository(db *mongo.Database) RentalRepository {
	return &rentalRepository{
		collection: db.Collection("rentals"),
	}
}

func (r *rentalRepository) AddRental(ctx context.Context, rental types.Rental) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, rental)
	if err != nil {
		log.Printf("Error inserting rental: %v", err)
		return err
	}
	return nil
}

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
