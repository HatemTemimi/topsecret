package service

import (
	"context"
	"errors"
	"server/internal/user/repository"
	"server/internal/user/types"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	CreateUser(ctx context.Context, user *types.User) (string, error)
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	UpdateUser(ctx context.Context, id string, updateData bson.M) error
	DeleteUser(ctx context.Context, id string) error
	Authenticate(ctx context.Context, email, password string) (*types.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser validates and creates a new user
func (s *userService) CreateUser(ctx context.Context, user *types.User) (string, error) {
	// Validate required fields
	if user.Email == "" || user.Password == "" {
		return "", errors.New("email and password are required")
	}

	// Check if the email is already in use
	existingUser, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		return "", errors.New("email is already registered")
	}

	// Delegate to the repository to create the user
	id, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

// GetUserByID retrieves a user by their ID
func (s *userService) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.FindUserByID(ctx, id)
}

// GetUserByEmail retrieves a user by their email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return s.repo.FindUserByEmail(ctx, email)
}

// UpdateUser updates a user's details
func (s *userService) UpdateUser(ctx context.Context, id string, updateData bson.M) error {
	if id == "" {
		return errors.New("id is required")
	}
	if len(updateData) == 0 {
		return errors.New("update data is required")
	}
	return s.repo.UpdateUser(ctx, id, updateData)
}

// DeleteUser deletes a user by their ID
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.repo.DeleteUser(ctx, id)
}

// Authenticate authenticates a user using their email and password
func (s *userService) Authenticate(ctx context.Context, email, password string) (*types.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.AuthenticateUser(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
