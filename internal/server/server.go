package server

import (
	"crypto/rsa"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"go-fiber-postgres-template/internal/database"

	"github.com/goccy/go-json"
)

type FiberServer struct {
	*fiber.App

	db database.Service

	key *rsa.PrivateKey

	validator *validator.Validate
}
type ErrorResponse struct {
	Details string `json:"details"`
}

func New(privateKey *rsa.PrivateKey) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-fiber-postgres-template",
			AppName:      "go-fiber-postgres-template",
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				var e *fiber.Error
				code := fiber.StatusInternalServerError
				errorResponse := ErrorResponse{
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

		db:        database.New(),
		key:       privateKey,
		validator: validator.New(),
	}
	return server
}
