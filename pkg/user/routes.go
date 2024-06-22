package user

import "github.com/gofiber/fiber/v2"

func UserRouters(app *fiber.App) {
	app.Get("/users", GetAllUsers)
	app.Post("/register", CreateUser)
	app.Get("/users/me", GetAllUsers)
	app.Put("/users/:id", GetAllUsers)
	app.Patch("/users/:id", GetAllUsers)
	app.Delete("/users/:id", GetAllUsers)
}
