package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user entities.User) (*entities.User, error)
	GetById(ctx context.Context, userId uuid.UUID) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, userId uuid.UUID) error
}
