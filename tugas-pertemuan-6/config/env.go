package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ProjectConfig struct {
	DatabaseUrl string
	AdminName   string
	AdminPass   string
	JWTSecret   string
}

var configInstance *ProjectConfig

// JWT will expire in a day
const JWT_EXPIRATION_LENGTH = 24 * 60 * 60

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return "", fmt.Errorf("tidak dapat menemukan %s di environment", key)
}

func initConfig() ProjectConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dbUrl, err := getEnv("DATABASE_URL")
	if err != nil {
		panic(err)
	}

	jwtSecret, err := getEnv("JWT_SECRET")
	if err != nil {
		panic(err)
	}

	adminName, err := getEnv("ADMIN_NAME")
	if err != nil {
		panic(err)
	}
	if len(adminName) < 3 {
		panic("ADMIN_NAME length is less than 3 characters")
	}

	adminPass, err := getEnv("ADMIN_PASS")
	if err != nil {
		panic(err)
	}
	if len(adminPass) < 8 {
		panic("ADMIN_PASS length is less than 8 characters")
	}

	configInstance = &ProjectConfig{
		DatabaseUrl: dbUrl,
		AdminName:   adminName,
		AdminPass:   adminPass,
		JWTSecret:   jwtSecret,
	}

	return *configInstance
}

func GetConfig() ProjectConfig {
	if configInstance == nil {
		return initConfig()
	}

	return *configInstance
}
