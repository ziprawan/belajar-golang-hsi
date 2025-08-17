package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	NIM      string `gorm:"not null" json:"nim"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	Major    string `gorm:"not null" json:"major"`
	Semester int    `gorm:"not null" json:"semester"`
}

type StudentModel struct {
	Student

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
