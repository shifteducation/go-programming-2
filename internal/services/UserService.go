package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/interfaces"
	"github.com/shifteducation/user-service/internal/models"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) Save(ctx context.Context, user models.User) models.User {
	return s.userRepository.Save(ctx, user)
}

func (s UserService) GetById(ctx context.Context, userId uuid.UUID) models.User {
	return s.userRepository.GetById(ctx, userId)
}

func (s UserService) GetAll(ctx context.Context) []models.User {
	return s.userRepository.GetAll(ctx)
}

func (s UserService) Update(ctx context.Context, user models.User) models.User {
	return s.userRepository.Update(ctx, user)
}

func (s UserService) Delete(ctx context.Context, userId uuid.UUID) {
	s.userRepository.Delete(ctx, userId)
}
