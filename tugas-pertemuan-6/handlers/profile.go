package handlers

import (
	"fmt"
	"pertemuan6/models"

	"github.com/gofiber/fiber/v2"
)

// GetProfile godoc
// @Summary Ambil profil
// @Description Ambil informasi profil pengguna berdasarkan kredensial yang diberikan di kredensial "Bearer"
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ProfileResponse "Profil pengguna"
// @Failure 401 {object} models.ErrorResponse "Kesalahan kredensial"
// @Failure 500 {object} models.ErrorResponse "Kesalahan internal server"
// @Router /profile [get]
func GetProfileHandler(c *fiber.Ctx) error {
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
