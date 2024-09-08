package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}]:${port} ${status} ${latency} - ${method} ${path} \n",
		TimeFormat: "2006-01-02T15:04:05",
	}))

	s.App.Get("/.well-known/jwks.json", s.JwksHandler)
	s.App.Get("/health", s.healthHandler)
	s.App.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "list",
		Filter: swagger.FilterConfig{
			Enabled: true,
		},
		CustomStyle: "body { margin: 0; }",
	}))

	app := s.App.Group("/api")
	v1 := app.Group("v1")

	{
		auth := v1.Group("auth")
		auth.Post("/register", s.RegisterHandler)
		auth.Post("/login", s.LoginHandler)
		auth.Post("/refresh-token", s.RefreshTokenHandler)
	}

	{
		users := v1.Group("users")
		users.Use(jwtMiddleware.New(jwtMiddleware.Config{
			SigningKey: jwtMiddleware.SigningKey{
				JWTAlg: jwtMiddleware.RS256,
				Key:    s.key.Public(),
			},
			SuccessHandler: func(c *fiber.Ctx) error {
				user := c.Locals("user").(*jwt.Token)
				claims := user.Claims.(jwt.MapClaims)

				if claims["typ"] != "access" {
					return fiber.NewError(fiber.StatusUnauthorized, "Invalid token type")
				}

				return c.Next()
			},
		}))
		users.Get("/me", s.GetMeHandler)
	}
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
