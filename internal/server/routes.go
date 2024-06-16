package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)
}

// /	 Hello World godoc
//
//	@Summary		Hello World example
//	@Description	Hello World
//	@Tags			Hello World
//	@Success		200	{object} map[string]string
//	@Failure		500	{object} map[string]string
//	@Router			/ [get]
func (*FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

// / Health godoc
//
//	@Summary		Health example
//	@Description	check server health
//	@Tags			Health
//	@Success		200	{object} map[string]string
//	@Failure		500	{object} map[string]string
//	@Router			/health [get]
func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
