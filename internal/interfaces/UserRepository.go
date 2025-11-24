package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/models"
)

type UserRepository interface {
	Save(ctx context.Context, user models.User) (*models.User, error)
	GetById(ctx context.Context, userId uuid.UUID) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, userId uuid.UUID) error
}
