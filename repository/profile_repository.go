package repository

import (
	"api_crud/entities"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CollectionProfile() ([]entities.Profile, error)

	DocumentProfile(id uint) (*entities.Profile, error)
}

type ProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{db: db}
}

func (p *ProfileRepositoryImpl) CollectionProfile() ([]entities.Profile, error) {
	var profiles []entities.Profile

	err := p.db.Find(&profiles).Error

	return profiles, err
}

func (p *ProfileRepositoryImpl) DocumentProfile(id uint) (*entities.Profile, error) {
	var profile entities.Profile

	if err := p.db.First(&profile, id).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
