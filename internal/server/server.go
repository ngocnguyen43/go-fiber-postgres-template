package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"go-fiber-postgres-template/internal/database"

	"github.com/goccy/go-json"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}
type ErrorResponse struct {
	Details string
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-fiber-postgres-template",
			AppName:      "go-fiber-postgres-template",
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				var code = fiber.StatusInternalServerError
				// Retrieve the custom status code if it's a *fiber.Error
				var e *fiber.Error
				var errorResponse = ErrorResponse{
					Details: "Internal Server Error",
				}
				if errors.As(err, &e) {
					code = e.Code
					errorResponse.Details = e.Message

				}
				if err != nil {
					// In case the SendFile fails
					return ctx.Status(code).JSON(errorResponse)
				}

				// Return from handler
				return nil
			},
		}),

		db: database.New(),
	}
	return server
}
