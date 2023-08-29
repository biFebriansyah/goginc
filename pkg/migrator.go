package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var pgurl string

type migrator struct {
	*migrate.Migrate
}

func init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")

	pgurl = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable search_path=%s", host, user, password, port, dbName, schema)
}

func NewMigrator() *migrator {
	if pgurl == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", pgurl)
	if err != nil {
		log.Fatalf("Error connect database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	return &migrator{m}
}

func (m *migrator) Ups() {
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalf("Error applying migrations: %v", err)
		}
	}

	log.Println("Migrations UP successfully")
}

func (m *migrator) Downs() {
	if err := m.Drop(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalf("Error applying migrations: %v", err)
		}
	}

	log.Println("Migrations DOWN successfully")
}
