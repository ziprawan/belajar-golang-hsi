package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func initDB() {
	conf := GetConfig()

	db, err := gorm.Open(postgres.Open(conf.DatabaseUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal membuka koneksi:", err)
		os.Exit(1)
	}

	dbInstance = db
}

func GetDB() *gorm.DB {
	if dbInstance == nil {
		initDB()
	}

	return dbInstance
}
