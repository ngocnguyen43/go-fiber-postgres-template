package server

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-postgres-template/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-fiber-postgres-template",
			AppName:      "go-fiber-postgres-template",
		}),

		db: database.New(),
	}

	return server
}
