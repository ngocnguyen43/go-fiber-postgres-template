package database

import (
	"context"
	"fmt"
	"go-fiber-postgres-template/internal/models"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/jackc/pgx/v5/stdlib"    //nolint:golint // no need to check this
	_ "github.com/joho/godotenv/autoload" //nolint:golint // no need to check this
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	GetInstance() *gorm.DB
}

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	env        = os.Getenv("ENV")
	dbInstance *service
)

const (
	dbWaitCount      = 1000
	dbConnectionPool = 40
)

func New() Service {
	var loggerConfig logger.Config

	if env == "dev" {
		loggerConfig = logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		}
	}

	var newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)

	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s",
		username,
		password,
		net.JoinHostPort(host, port),
		database,
		schema,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	_, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	dbClient, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("db down: %v", err) // Log the error and terminate the program //nolint:gocritic // no need to check this
		return stats
	}

	// Ping the database with context
	if err = dbClient.Ping(); err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db ping error: %v", err)
		log.Printf("db ping error: %v", err) // Log the error
		return stats
	}
	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := dbClient.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > dbConnectionPool { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > dbWaitCount {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern." //nolint:lll //keep this line
	}
	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	postgresDB, _ := s.db.DB()
	return postgresDB.Close()
}

func (s *service) GetInstance() *gorm.DB {
	return s.db
}

func GetTestInstance() Service {
	testDB, err := gorm.Open(sqlite.Open("gorm_test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("open test db error: %v", err)
	}
	err = testDB.AutoMigrate(&models.User{}, &models.RefreshToken{}, &models.RefreshTokenFamily{})
	if err != nil {
		log.Printf("migrate test db error: %v", err)
	}
	dbInstance = &service{
		db: testDB,
	}
	return dbInstance
}

func CloseTestDBInstance(testDB *gorm.DB) error {
	db, err := testDB.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	err = os.Remove("gorm_test.db")
	return err
}
