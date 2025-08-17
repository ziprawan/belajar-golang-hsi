package models

import "pertemuan6/config"

func MigrateAll() error {
	db := config.GetDB()

	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&StudentModel{})
	if err != nil {
		return err
	}

	return nil
}
