package models

import "gorm.io/gorm"

type TugasModel struct {
	gorm.Model

	MahasiswaID uint   `gorm:"uniqueIndex:idx_mhsid_judul;not null"`
	Judul       string `gorm:"uniqueIndex:idx_mhsid_judul;not null"`
	Deskripsi   string `gorm:"not null"`

	Mahasiswa MahasiswaModel `gorm:"foreignKey:MahasiswaID"`
}
