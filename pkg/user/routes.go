package user

import "github.com/gofiber/fiber/v2"

func UserRouters(app *fiber.App) {
	app.Get("/users", GetAllUsers)
}
