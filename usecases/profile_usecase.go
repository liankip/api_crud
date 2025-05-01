package usecases

import (
	"api_crud/entities"
	"api_crud/repository"
)

type ProfileUsecase struct {
	ProfileRepository repository.ProfileRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{
		ProfileRepository: profileRepository,
	}
}

func (profileUseCase *ProfileUsecase) CollectionProfile() ([]entities.Profile, error) {
	profile, err := profileUseCase.ProfileRepository.CollectionProfile()

	return profile, err
}

func (profileUseCase *ProfileUsecase) DocumentProfile(id uint) (*entities.Profile, error) {
	profile, err := profileUseCase.ProfileRepository.DocumentProfile(id)

	return profile, err
}

func (profileUseCase *ProfileUsecase) CreateProfile(createProfile entities.CreateProfile) (*entities.Profile, error) {
	profile, err := profileUseCase.ProfileRepository.CreateProfile(createProfile)

	return profile, err
}

func (profileUseCase *ProfileUsecase) UpdateProfile(updateProfile entities.UpdateProfile) (*entities.Profile, error) {
	profile, err := profileUseCase.ProfileRepository.UpdateProfile(updateProfile)

	return profile, err
}

func (profileUseCase *ProfileUsecase) DeleteProfile(userID uint) error {
	err := profileUseCase.ProfileRepository.DeleteProfile(userID)

	return err
}
