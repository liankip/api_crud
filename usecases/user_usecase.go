package usecases

import "api_crud/repository"

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}
