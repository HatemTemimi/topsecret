// repository/rental_repository.go
package repository

import (
	"context"
	"log"
	"time"

	"server/internal/rental/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type RentalRepository interface {
	AddRental(ctx context.Context, rental types.Rental) error
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
