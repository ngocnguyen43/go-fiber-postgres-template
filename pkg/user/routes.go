package user

import "github.com/gofiber/fiber/v2"

func UserRouters(app fiber.Router) {
	app.Get("/users", GetAllUsers)
	app.Get("/users/me", GetAllUsers)
	app.Put("/users/:id", GetAllUsers)
	app.Patch("/users/:id", GetAllUsers)
	app.Delete("/users/:id", GetAllUsers)
}
