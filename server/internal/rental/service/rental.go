// service/rental_service.go
package service

import (
	"context"
	"server/internal/rental/repository"
	"server/internal/rental/types"
)

type RentalService interface {
	AddRental(ctx context.Context, rental types.Rental) error
	GetAllRentals(ctx context.Context) ([]types.Rental, error)
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(repo repository.RentalRepository) RentalService {
	return &rentalService{repo: repo}
}

func (s *rentalService) AddRental(ctx context.Context, rental types.Rental) error {
	// Perform any business logic checks, if needed
	return s.repo.AddRental(ctx, rental)
}

// GetAllRentals retrieves all rentals by calling the repository layer
func (s *rentalService) GetAllRentals(ctx context.Context) ([]types.Rental, error) {
	return s.repo.GetAllRentals(ctx)
}
