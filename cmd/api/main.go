package main

import (
	"fmt"
	internalServer "go-fiber-postgres-template/internal/server"
	"os"
	"strconv"

	"github.com/gofiber/swagger"

	_ "go-fiber-postgres-template/docs"

	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

//	@title						Fiber Example API
//	@version					1.0
//	@description				This is a sample swagger for Fiber
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.email				fiber@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8080
//	@BasePath					/api
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization

func main() {
	server := internalServer.New()
	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} ${latency} - ${method} ${path} \n",
	}))
	server.RegisterFiberRoutes()
	server.Get("/docs/*", swagger.HandlerDefault) // default

	server.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
