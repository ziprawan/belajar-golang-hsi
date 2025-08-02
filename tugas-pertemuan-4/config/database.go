package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return "", fmt.Errorf("tidak dapat menemukan %s di environment", key)
}

func getPostgresDSN() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	host, err := getEnv("DB_HOST")
	if err != nil {
		return "", err
	}
	user, err := getEnv("DB_USER")
	if err != nil {
		return "", err
	}
	pass, err := getEnv("DB_PASS")
	if err != nil {
		return "", err
	}
	port, err := getEnv("DB_PORT")
	if err != nil {
		return "", err
	}
	db_name, err := getEnv("DB_NAME")
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		host, user, pass, port, db_name,
	)
	return dsn, nil
}

func initDB() {
	dsn, err := getPostgresDSN()
	if err != nil {
		fmt.Println("Gagal menginisiasi koneksi:", err)
		os.Exit(1)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
