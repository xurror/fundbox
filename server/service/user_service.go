package service

import (
	"context"
	"getting-to-go/model"

	"github.com/google/uuid"
)

// UserService provides user-related services
type UserService interface {
	CreateUser(ctx context.Context, auth0Id string) (*model.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserByAuth0Id(ctx context.Context, auth0Id string) (*model.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUsers(ctx context.Context, limit int, offset interface{}) ([]*model.User, interface{}, error)

	// GetUserContributions(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.Contribution, error)
}
