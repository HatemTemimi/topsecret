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
	GetRentalsByUserID(ctx context.Context, userID string) ([]types.Rental, error) // New Method
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
	if rental.Name == "" {
		return errors.New("rental name cannot be empty")
	}
	if rental.StreetNumber == "" || rental.Street == "" || rental.City == "" || rental.Country == "" {
		return errors.New("address fields cannot be empty")
	}
	return s.repo.AddRental(ctx, rental)
}

// GetAllRentals retrieves all rentals
func (s *rentalService) GetAllRentals(ctx context.Context) ([]types.Rental, error) {
	return s.repo.GetAllRentals(ctx)
}

// GetRentalByID retrieves a single rental by its ID
func (s *rentalService) GetRentalByID(ctx context.Context, id string) (*types.Rental, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.GetRentalByID(ctx, id)
}

// GetRentalsByUserID retrieves all rentals for a specific user
func (s *rentalService) GetRentalsByUserID(ctx context.Context, userID string) ([]types.Rental, error) {
	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}

	return s.repo.GetRentalsByUserID(ctx, userID)
}

// UpdateRental updates an existing rental
func (s *rentalService) UpdateRental(ctx context.Context, id string, updatedData types.Rental) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	if updatedData.Name == "" {
		return errors.New("rental name cannot be empty")
	}
	if updatedData.StreetNumber == "" || updatedData.Street == "" || updatedData.City == "" || updatedData.Country == "" {
		return errors.New("address fields cannot be empty")
	}
	return s.repo.UpdateRental(ctx, id, updatedData)
}

// DeleteRental deletes a rental by its ID
func (s *rentalService) DeleteRental(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.DeleteRental(ctx, id)
}
