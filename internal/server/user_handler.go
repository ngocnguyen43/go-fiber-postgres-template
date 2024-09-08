package server

import (
	"go-fiber-postgres-template/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Get me godoc
// @Summary 	Get my infomation information
// @Schemes
// @Description This endpoint allows a user to get their information.
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} models.User
// @Router 		/users/me [get]
// @Security 	JWT
func (s *FiberServer) GetMeHandler(c *fiber.Ctx) error {
	userData := c.Locals("user").(*jwt.Token)
	claims := userData.Claims.(jwt.MapClaims)
	userID := claims["sub"].(float64)
	var user models.User
	if err := s.db.GetInstance().Where("id = ?", userID).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
