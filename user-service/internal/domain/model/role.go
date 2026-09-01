package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        int64
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}

func (Role) TableName() string {
	return "roles"
}
