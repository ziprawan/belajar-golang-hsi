package models

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	err := db.AutoMigrate(&HasilModel{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&TugasModel{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&MahasiswaModel{})
	if err != nil {
		return err
	}

	return nil
}
