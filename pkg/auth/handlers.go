package auth

import (
	"crypto/rand"
	"go-fiber-postgres-template/pkg/user"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2"
)

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Register a new user
//	@Tags			Auth
//	@Produce		json
//	@Param			data	body	RegisterRequest	true	"The input user struct"
//	@Success		200	{object} user.User
//	@Router			/register [post]
func Register(c *fiber.Ctx) error {
	// db := database.New().GetInstance()

	user := new(user.User)

	register := &RegisterRequest{}

	if err := register.bind(c, user); err != nil {
		return err
	}
	// salt, _ := generateSalt(10)
	// hash := hashPasswordPbkd2("mninhngocnguyen", salt)
	// fmt.Println(verifyPassword(hash, "123"))
	// err := db.Create(&user).Error
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	return c.JSON(user)
}
