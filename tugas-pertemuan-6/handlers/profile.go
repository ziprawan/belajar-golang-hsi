package handlers

import (
	"fmt"
	"pertemuan6/models"

	"github.com/gofiber/fiber/v3"
)

func GetProfileHandler(c fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		fmt.Println("Invalid locals type for \"user\"")
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	return c.JSON(models.ProfileResponse{UserSafe: models.UserSafe{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}})
}
