package main

import (
	"fmt"
	"go-fiber-postgres-template/internal/server"
	"go-fiber-postgres-template/internal/utils"
	"log"
	"os"
	"strconv"

	_ "go-fiber-postgres-template/docs"

	_ "github.com/joho/godotenv/autoload"
)

// @title						Fiber Example API
// @version						3.0.0
// @description					This is a sample swagger for Fiber
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.email				fiber@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
func main() {
	keysDir := "keys"
	bits := 2048

	if err := os.MkdirAll(keysDir, 0700); err != nil {
		log.Printf("Failed to create keys directory: %v\n", err)
		return
	}
	privateKey, err := utils.EnsureRSAKeys(keysDir, bits)
	if err != nil {
		log.Printf("Failed to ensure RSA keys: %v\n", err)
		return
	}
	server := server.New(privateKey)
	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err = server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}
