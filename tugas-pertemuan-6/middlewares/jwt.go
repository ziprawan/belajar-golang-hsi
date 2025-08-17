package middlewares

import (
	"fmt"
	"pertemuan6/config"
	"pertemuan6/models"
	"pertemuan6/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(c *fiber.Ctx) error {
	conf := config.GetConfig()

	auth := c.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return c.Status(401).JSON(models.UnauthorizedErrorResponse)
	}

	token := strings.Replace(auth, "Bearer ", "", 1)

	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method is not HMAC")
		}
		return []byte(conf.JWTSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(401).JSON(models.ErrorResponse{Error: err.Error()})
	}

	if !parsed.Valid {
		return c.Status(401).JSON(models.UnauthorizedErrorResponse)
	}

	sub, err := parsed.Claims.GetSubject()
	if err != nil {
		fmt.Println(err)
		return c.Status(401).JSON(models.UnauthorizedErrorResponse)
	}

	db := config.GetDB()

	user := models.UserModel{User: models.User{
		ID: utils.SafeParseUint(sub),
	}}
	tx := db.First(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		fmt.Println(tx.Error)
		return c.Status(401).JSON(models.UnauthorizedErrorResponse)
	}

	c.Locals("user", user.User)

	return c.Next()
}
