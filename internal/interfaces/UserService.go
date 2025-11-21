package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/models"
)

type UserService interface {
	Save(ctx context.Context, user models.User) models.User
	GetById(ctx context.Context, userId uuid.UUID) models.User
	GetAll(ctx context.Context) []models.User
	Update(ctx context.Context, user models.User) models.User
	Delete(ctx context.Context, userId uuid.UUID)
}
