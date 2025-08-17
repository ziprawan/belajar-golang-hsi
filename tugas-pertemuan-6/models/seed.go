package models

import (
	"pertemuan6/config"

	"github.com/matthewhartstonge/argon2"
)

func SeedUserAdmin() error {
	db := config.GetDB()
	conf := config.GetConfig()
	argon := argon2.DefaultConfig()

	var count int64
	db.Model(&UserModel{}).Count(&count)

	if count != 0 {
		return nil
	}

	encoded, err := argon.HashEncoded([]byte(conf.AdminPass))
	if err != nil {
		return err
	}

	user := UserModel{
		User: User{
			Role:     "admin",
			Username: conf.AdminName,
			Password: string(encoded),
		},
	}
	db.Create(&user)

	return nil
}
