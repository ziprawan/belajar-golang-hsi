package middlewares

import (
	"fmt"
	"pertemuan6/models"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		fmt.Println("Admin middleware should come after jwt middleware! Please re-check the middleware flows!")
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	if user.Role != "admin" {
		return c.Status(403).JSON(models.ForbiddenErrorResponse)
	}

	return c.Next()
}
