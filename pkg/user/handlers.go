package user

import (
	"go-fiber-postgres-template/internal/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.New().GetInstance()
	var users []User
	db.Find(&users)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": users})
}
