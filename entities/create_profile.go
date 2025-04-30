package entities

type CreateProfile struct {
	UserID    uint
	Bio       string `validate:"required"`
	AvatarUrl string `validate:"required"`
}
