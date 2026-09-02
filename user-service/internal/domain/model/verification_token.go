package model

import (
	"time"
)

type VerificationToken struct {
	ID        int64
	UserID    int64
	HashToken string
	TokenType string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time

	User User `gorm:"foreignKey:UserID;references:ID;OnDelete:RESTRICT"`
}
