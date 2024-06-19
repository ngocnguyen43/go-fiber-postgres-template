package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func getMigrationsPath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the absolute path to the migrations directory
	absolutePath := strings.TrimLeft("internal/database/migrations", string(os.PathSeparator))
	migrationsPath := filepath.Join(workingDir, absolutePath)
	absoluteMigrationsPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		log.Fatal(err)
	}
	return absoluteMigrationsPath
}
func main() {
	var exec string
	flag.StringVar(&exec, "exec", "up", "")
	flag.Parse()
	if exec != "up" && exec != "down" {
		log.Fatal("Error when exec migrate command")
	}
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	migrationsPath := getMigrationsPath()
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		database, driver)
	if err != nil {
		log.Fatal(err)
	}
	if exec == "up" {
		if err = m.Up(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err = m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}
