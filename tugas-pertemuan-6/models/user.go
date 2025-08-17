package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"` // Di real app, harus di-hash
	Role     string `gorm:"not null" json:"role"`     // "admin" atau "student"
}

type UserSafe struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserModel struct {
	User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
