package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/models"
	"log"
)

type UserPostgresRepository struct {
}

func (r UserPostgresRepository) Save(ctx context.Context, user models.User) (*models.User, error) {
	log.Print("Not implemented")
	return nil, nil
}

func (r UserPostgresRepository) GetById(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	log.Print("Not implemented")
	return nil, nil
}

func (r UserPostgresRepository) GetAll(ctx context.Context) ([]models.User, error) {
	log.Print("Not implemented")
	return make([]models.User, 0), nil
}

func (r UserPostgresRepository) Update(ctx context.Context, user models.User) (*models.User, error) {
	log.Print("Not implemented")
	return nil, nil
}

func (r UserPostgresRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	log.Print("Not implemented")
	return nil
}
