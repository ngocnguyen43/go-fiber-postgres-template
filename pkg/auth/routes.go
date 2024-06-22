package auth

import "github.com/gofiber/fiber/v2"

func AuthRouters(app fiber.Router) {
	app.Post("/register", Register)
}
