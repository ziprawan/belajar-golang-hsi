package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"pertemuan6/config"
	"pertemuan6/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Login godoc
// @Summary Masuk pengguna
// @Description Verifikasi pengguna dengan kredensial masuk ("username" dan "password") yang sudah diberikan, lalu mengembalikan info pengguna dan token JWT yang bisa digunakan di header "Authorization"
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.AuthRequest true "Kredensial masuk"
// @Success 200 {object} models.LoginResponse "Masuk berhasil"
// @Failure 400 {object} models.ErrorResponse "Isian permintaan salah"
// @Failure 401 {object} models.ErrorResponse "Kesalahan kredensial"
// @Failure 500 {object} models.ErrorResponse "Kesalahan internal server"
// @Router /auth/login [post]
func LoginHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	conf := config.GetConfig()
	body := c.Body()

	req := models.AuthRequest{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	if len(req.Username) < 3 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "username field length is less than 3"})
	}
	if len(req.Password) < 8 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "password field length is less than 8"})
	}

	user := models.UserModel{User: models.User{}}
	tx := db.First(&user, "username = ?", req.Username)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(models.ErrorResponse{Error: "Username not found"})
		}

		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	verified, err := argon2.VerifyEncoded([]byte(req.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	if !verified {
		return c.Status(403).JSON(models.ErrorResponse{Error: "Invalid password"})
	}

	iat := time.Now().Unix()
	nbf := iat
	exp := nbf + config.JWT_EXPIRATION_LENGTH
	sub := strconv.FormatUint(uint64(user.ID), 10)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": iat,
		"nbf": nbf,
		"exp": exp,
		"sub": sub,
	})
	token, err := t.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	userSafe := models.UserSafe{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	return c.JSON(models.LoginResponse{User: userSafe, Token: token})
}

// Register godoc
// @Summary Mendaftar pengguna baru
// @Description Mendaftar pengguna baru dengan role/status "student" kemudian mengembalikan informasi pengguna yang sudah dibuat
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.AuthRequest true "Kredensial pengguna baru"
// @Success 201 {object} models.RegisterResponse "Informasi pengguna yang baru didaftarkan"
// @Failure 400 {object} models.ErrorResponse "Isian permintaan salah"
// @Failure 403 {object} models.ErrorResponse "\"username\" sudah diambil"
// @Failure 500 {object} models.ErrorResponse "Kesalahan internal server"
// @Router /auth/register [post]
func RegisterHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	body := c.Body()

	req := models.AuthRequest{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	if len(req.Username) < 3 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "username field length is less than 3"})
	}
	if len(req.Password) < 8 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "password field length is less than 8"})
	}

	argon := argon2.DefaultConfig()
	password, err := argon.HashEncoded([]byte(req.Password))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	user := models.UserModel{User: models.User{
		Username: req.Username,
		Password: string(password),
		Role:     "student",
	}}
	tx := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}
	if tx.RowsAffected == 0 {
		return c.Status(403).JSON(models.ErrorResponse{Error: "Username taken"})
	}

	return c.Status(201).JSON(models.RegisterResponse{UserSafe: models.UserSafe{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}})
}
