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

func (userUseCase *UserUsecase) Signin(signin entities.Signin) (*entities.User, error) {
	user, err := userUseCase.UserRepository.FindUserByEmail(&signin.Email)

	if err != nil || !utils.CheckPasswordHash(signin.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (userUseCase *UserUsecase) Signup(signup entities.Signup) error {
	_, err := userUseCase.UserRepository.FindUserByEmail(signup.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	password, _ := utils.HashPassword(signup.Password)

	user := &entities.User{
		Username: signup.Username,
		Email:    signup.Email,
		Password: password,
	}

	return userUseCase.UserRepository.CreateUser(user)
}
