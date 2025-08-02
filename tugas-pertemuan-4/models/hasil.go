package models

import "gorm.io/gorm"

type HasilModel struct {
	gorm.Model

	TugasID uint `gorm:"not null"`
	Nilai   uint `gorm:"not null"`

	Tugas TugasModel `gorm:"foreignKey:TugasID"`
}
