package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/custom_errors"
	"github.com/shifteducation/user-service/internal/dto"
	"github.com/shifteducation/user-service/internal/entities"
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
			Id:        uuid.New(),
			City:      userDto.Address.City,
			Street:    userDto.Address.Street,
			Building:  userDto.Address.Building,
			Apartment: userDto.Address.Apartment,
		}
	}

	user := models.User{
		Id:        uuid.New(),
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Age:       userDto.Age,
		Address:   address,
	}

	userEntity, err := s.userRepository.Save(ctx, s.userToUserEntity(user))
	if err != nil {
		return nil, err
	}

	savedUser := s.userEntityToUser(*userEntity)
	return &savedUser, nil
}

func (s UserService) GetById(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	userEntity, err := s.userRepository.GetById(ctx, userId)
	if err == nil && userEntity == nil {
		return nil, custom_errors.UserNotFoundError{}
	} else if userEntity != nil {
		user := s.userEntityToUser(*userEntity)
		return &user, err
	}

	return nil, err
}

func (s UserService) GetAll(ctx context.Context) ([]models.User, error) {
	userEntities, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for _, userEntity := range userEntities {
		users = append(users, s.userEntityToUser(userEntity))
	}
	return users, nil
}

func (s UserService) Update(ctx context.Context, user models.User) error {
	return s.userRepository.Update(ctx, s.userToUserEntity(user))
}

func (s UserService) Delete(ctx context.Context, userId uuid.UUID) error {
	return s.userRepository.Delete(ctx, userId)
}

func (s UserService) userToUserEntity(user models.User) entities.User {
	var addressEntityP *entities.Address
	if user.Address == nil {
		addressEntityP = nil
	} else {
		addressEntity := s.addressToAddressEntity(*user.Address, user.Id)
		addressEntityP = &addressEntity
	}

	return entities.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Address:   addressEntityP,
	}
}

func (s UserService) userEntityToUser(userEntity entities.User) models.User {
	var addressP *models.Address
	if userEntity.Address == nil {
		addressP = nil
	} else {
		address := s.addressEntityToAddress(*userEntity.Address)
		addressP = &address
	}

	return models.User{
		Id:        userEntity.Id,
		FirstName: userEntity.FirstName,
		LastName:  userEntity.LastName,
		Age:       userEntity.Age,
		Address:   addressP,
	}
}

func (s UserService) addressToAddressEntity(address models.Address, userId uuid.UUID) entities.Address {
	return entities.Address{
		Id:        address.Id,
		City:      address.City,
		Street:    address.Street,
		Building:  address.Building,
		Apartment: address.Apartment,
		UserId:    &userId,
	}
}

func (s UserService) addressEntityToAddress(addressEntity entities.Address) models.Address {
	return models.Address{
		Id:        addressEntity.Id,
		City:      addressEntity.City,
		Street:    addressEntity.Street,
		Building:  addressEntity.Building,
		Apartment: addressEntity.Apartment,
	}
}
