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

	HasAccess(userID int, accessName string) bool
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

func (repository *UserRepositoryImpl) HasAccess(userID int, accessName string) bool {
	var count int
	query := `SELECT COUNT(*)
    FROM users u
    JOIN user_role ur ON u.id = ur.user_id
    JOIN role_access ra ON ur.role_id = ra.role_id
    JOIN access a ON ra.access_id = a.access_id
    WHERE u.id = $1 AND a.access_name = $2`

	err := repository.db.Raw(query, userID, accessName).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
