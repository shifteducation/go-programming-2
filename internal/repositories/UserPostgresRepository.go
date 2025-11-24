package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserPostgresRepository {
	return UserPostgresRepository{
		db: db,
	}
}

func (r UserPostgresRepository) Save(ctx context.Context, user entities.User) (*entities.User, error) {
	err := gorm.G[entities.User](r.db).Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserPostgresRepository) GetById(ctx context.Context, userId uuid.UUID) (*entities.User, error) {
	user, err := gorm.G[entities.User](r.db).
		Where("id = ?", userId).
		Preload("Address", nil). //todo join
		First(ctx)
	return &user, err
}

func (r UserPostgresRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	users, err := gorm.G[entities.User](r.db).
		Joins(clause.LeftJoin.Association("Address"), nil).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserPostgresRepository) Update(ctx context.Context, user entities.User) error {
	log.Print("Not implemented")
	return nil
}

func (r UserPostgresRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	log.Print("Not implemented")
	return nil
}
