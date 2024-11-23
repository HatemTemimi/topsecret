package service

import (
	"context"
	"errors"
	"server/internal/rental/repository"
	"server/internal/rental/types"
)

type RentalService interface {
	AddRental(ctx context.Context, rental types.Rental) error
	GetAllRentals(ctx context.Context) ([]types.Rental, error)
	GetRentalByID(ctx context.Context, id string) (*types.Rental, error)
	UpdateRental(ctx context.Context, id string, updatedData types.Rental) error
	DeleteRental(ctx context.Context, id string) error
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(repo repository.RentalRepository) RentalService {
	return &rentalService{repo: repo}
}

// AddRental validates and adds a new rental
func (s *rentalService) AddRental(ctx context.Context, rental types.Rental) error {
	// Perform basic validation
	if rental.Name == "" {
		return errors.New("rental name cannot be empty")
	}
	if rental.StreetNumber == "" || rental.Street == "" || rental.City == "" || rental.Country == "" {
		return errors.New("address fields cannot be empty")
	}

	// Delegate to the repository layer
	return s.repo.AddRental(ctx, rental)
}

// GetAllRentals retrieves all rentals
func (s *rentalService) GetAllRentals(ctx context.Context) ([]types.Rental, error) {
	return s.repo.GetAllRentals(ctx)
}

// GetRentalByID retrieves a single rental by its ID
func (s *rentalService) GetRentalByID(ctx context.Context, id string) (*types.Rental, error) {
	// Validate ID
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	// Delegate to the repository layer
	return s.repo.GetRentalByID(ctx, id)
}

// UpdateRental updates an existing rental
func (s *rentalService) UpdateRental(ctx context.Context, id string, updatedData types.Rental) error {
	// Validate ID
	if id == "" {
		return errors.New("id cannot be empty")
	}

	// Validate update data
	if updatedData.Name == "" {
		return errors.New("rental name cannot be empty")
	}
	if updatedData.StreetNumber == "" || updatedData.Street == "" || updatedData.City == "" || updatedData.Country == "" {
		return errors.New("address fields cannot be empty")
	}

	// Delegate to the repository layer
	return s.repo.UpdateRental(ctx, id, updatedData)
}

// DeleteRental deletes a rental by its ID
func (s *rentalService) DeleteRental(ctx context.Context, id string) error {
	// Validate ID
	if id == "" {
		return errors.New("id cannot be empty")
	}

	// Delegate to the repository layer
	return s.repo.DeleteRental(ctx, id)
}
