package user

import (
	"go-fiber-postgres-template/pkg/core"

	"github.com/gofiber/fiber/v2"
)

func UserRouters(app fiber.Router) {
	app.Get("/users", core.IsAuthenticated(), GetAllUsers)
	app.Get("/users/me", core.IsAuthenticated(), GetAllUsers)
	app.Put("/users/:id", GetAllUsers)
	app.Patch("/users/:id", GetAllUsers)
	app.Delete("/users/:id", GetAllUsers)
}
