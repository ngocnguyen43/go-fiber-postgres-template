package main

import (
	"fmt"
	"go-fiber-postgres-template/internal/models"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.RefreshToken{}, &models.RefreshTokenFamily{}, &models.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	if _, err = io.WriteString(os.Stdout, stmts); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to stdout: %v\n", err)
	}
}
