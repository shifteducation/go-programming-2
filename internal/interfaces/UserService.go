package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/dto"
	"github.com/shifteducation/user-service/internal/models"
)

type UserService interface {
	Create(ctx context.Context, userDto dto.CreateUserRequest) (*models.User, error)
	GetById(ctx context.Context, userId uuid.UUID) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	UpdateAdrr(ctx context.Context, userId uuid.UUID, adrDTO dto.UpdateAddressRequest) error
	Delete(ctx context.Context, userId uuid.UUID) error
}
