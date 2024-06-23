package core

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: os.Getenv("APP_SECRET")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == jwtware.ErrJWTMissingOrMalformed.Error() {
				return fiber.NewError(fiber.StatusBadRequest, "Missing or Malformed JWT")
			}
			return c.Next()
		},
	})
}
