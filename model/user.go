package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID uint `gorm:"primaryKey"`

	CreatedAt time.Time `gorm:"autoCreateTime"` // Use autoCreateTime
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Use autoUpdateTime

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
}

func (User) TableName() string {
	return "user" // Explicitly set the table name to "user"  otherwise gorm pluralizes it always...
}
