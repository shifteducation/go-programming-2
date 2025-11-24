package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/custom_errors"
	"github.com/shifteducation/user-service/internal/dto"
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

func (s UserService) Create(ctx context.Context, userDto dto.CreateUserRequest) (*models.User, error) {
	var address *models.Address

	if userDto.Address != nil {
		address = &models.Address{
			Id:        uuid.Nil,
			City:      userDto.Address.City,
			Street:    userDto.Address.Street,
			Building:  userDto.Address.Building,
			Apartment: userDto.Address.Apartment,
		}
	}

	user := models.User{
		Id:        uuid.Nil,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Age:       userDto.Age,
		Address:   address,
	}

	return s.userRepository.Save(ctx, user)
}

func (s UserService) GetById(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetById(ctx, userId)
	if err == nil && user == nil {
		return nil, custom_errors.UserNotFoundError{}
	}
	return user, err
}

func (s UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s UserService) Update(ctx context.Context, user models.User) error {
	return s.userRepository.Update(ctx, user)
}

func (s UserService) Delete(ctx context.Context, userId uuid.UUID) error {
	return s.userRepository.Delete(ctx, userId)
}
