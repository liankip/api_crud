package repository

import (
	"api_crud/entities"
	"errors"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	CollectionProfile() ([]entities.Profile, error)

	DocumentProfile(id uint) (*entities.Profile, error)

	CreateProfile(profile entities.CreateProfile) (*entities.Profile, error)

	UpdateProfile(profile entities.UpdateProfile) (*entities.Profile, error)
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

func (p *ProfileRepositoryImpl) CreateProfile(profile entities.CreateProfile) (*entities.Profile, error) {
	var createProfile entities.Profile

	err := p.db.Where("user_id = ?", profile.UserID).First(&createProfile).Error
	if err != nil {
		return nil, errors.New("profile already exists")
	}

	createProfile.UserID = profile.UserID
	createProfile.Bio = profile.Bio
	createProfile.AvatarURL = profile.AvatarUrl

	if err := p.db.Create(&createProfile).Error; err != nil {
		return nil, err
	}

	return &createProfile, nil
}

func (p *ProfileRepositoryImpl) UpdateProfile(profile entities.UpdateProfile) (*entities.Profile, error) {
	var updateProfile entities.Profile

	err := p.db.Where("user_id = ?", profile.UserID).First(&updateProfile).Error
	if err != nil {
		return nil, errors.New("profile doesnt exists")
	}

	updateProfile.Bio = profile.Bio
	updateProfile.AvatarURL = profile.AvatarUrl

	if err := p.db.Save(&updateProfile).Error; err != nil {
		return nil, err
	}

	return &updateProfile, nil
}
