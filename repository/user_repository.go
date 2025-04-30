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

	HasAccess(userID uint, accessName string) bool
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

func (repository *UserRepositoryImpl) HasAccess(userID uint, accessName string) bool {
	var count int64

	err := repository.db.
		Table("users").
		Joins("JOIN user_role ON users.id = user_role.user_id").
		Joins("JOIN role_access ON user_role.role_id = role_access.role_id").
		Joins("JOIN access ON role_access.access_id = access.access_id").
		Where("users.id = ?", userID).
		Where("access.access_name = ?", accessName).
		Count(&count).Error

	if err != nil {
		return false
	}

	return count > 0
}
