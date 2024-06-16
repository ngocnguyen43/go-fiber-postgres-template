package main

import (
	"fmt"
	internalServer "go-fiber-postgres-template/internal/server"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := internalServer.New()

	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
