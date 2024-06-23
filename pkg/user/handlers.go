package user

import (
	"go-fiber-postgres-template/internal/database"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
//
//	@Summary		Get All Users
//	@Description	Retrieves a list of all users from the database
//	@Tags			Users
//	@Produce		json
//	@Success		200	{array} User
//	@Router			/users [get]
//
// @Security JWT
func GetAllUsers(c *fiber.Ctx) error {
	db := database.New().GetInstance()
	var users []User
	db.Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

// CreateUser godoc
//
//	@Summary		Create a New User
//	@Description	Create a new user in the database
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			data	body	User	true	"The input user struct"
//	@Success		200		{object}	User			"ok"
//	@Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.New().GetInstance()

	user := new(User)

	err := c.BodyParser(user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	err = db.Create(&user).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(user)
}
