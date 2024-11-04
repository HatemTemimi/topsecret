// service/rental_service.go
package service

import (
	"context"
	"server/internal/rental/repository"
	"server/internal/rental/types"
)

type RentalService interface {
	AddRental(ctx context.Context, rental types.Rental) error
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
