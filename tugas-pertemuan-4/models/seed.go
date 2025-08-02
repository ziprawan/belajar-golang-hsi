package models

import "gorm.io/gorm"

func SeedMahasiswa(db *gorm.DB) {
	mhs := []MahasiswaModel{
		{Nama: "Andi"},
		{Nama: "Budi"},
		{Nama: "Cecep"},
		{Nama: "Denis"},
		{Nama: "Erlangga"},
	}

	db.Create(&mhs)
}
