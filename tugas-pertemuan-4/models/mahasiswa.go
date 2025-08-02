package models

import "gorm.io/gorm"

type MahasiswaModel struct {
	gorm.Model

	Nama string `gorm:"not null"`

	Tugas []TugasModel `gorm:"foreignKey:MahasiswaID"`
}
