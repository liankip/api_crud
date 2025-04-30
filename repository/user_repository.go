package repository

import (
	"api_crud/entities"
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email *string) (*entities.User, error)

	CreateUser(user *entities.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repository *UserRepositoryImpl) FindUserByEmail(email *string) (*entities.User, error) {
	user := &entities.User{}

	result := repository.db.First(user, "email = ?", email)

	if errors.Is(result.Error, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	return user, result.Error
}

func (repository *UserRepositoryImpl) CreateUser(user *entities.User) error {
	result := repository.db.Create(user)

	if errors.Is(result.Error, sql.ErrNoRows) {
		return errors.New("user already exists")
	}

	return result.Error
}
