package usecases

import (
	"api_crud/entities"
	"api_crud/repository"
	"api_crud/utils"
	"errors"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (userUseCase *UserUsecase) Signin(email *string, password string) (*entities.User, error) {
	user, err := userUseCase.UserRepository.FindUserByEmail(email)

	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
