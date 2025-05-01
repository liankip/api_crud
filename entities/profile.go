package entities

import "time"

type Profile struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Bio       string    `gorm:"type:text"`
	AvatarURL string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
