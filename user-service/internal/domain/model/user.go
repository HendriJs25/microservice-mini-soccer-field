package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           int64
	UUID         uuid.UUID
	Name         string
	Username     string
	PasswordHash string
	Email        string
	IsVerified   bool
	RoleID       int64
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    gorm.DeletedAt

	Role Role `gorm:"foreignKey:RoleID;references:ID;OnDelete:RESTRICT"`
}

func (User) TableName() string {
	return "users"
}
