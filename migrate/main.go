package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/NayanPahuja/fam-bcknd-test/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Configure PostgreSQL connection
	db, err := sql.Open("postgres", "user="+config.Envs.DBUser+
		" password="+config.Envs.DBPassword+
		" host="+config.Envs.DBHost+
		" dbname="+config.Envs.DBName+
		" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a PostgreSQL driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrate/migrations", // Path to migration files
		"postgres",                  // Database name
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Read the command argument (up/down)
	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		// Apply migrations
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully!")
	}

	if cmd == "down" {
		// Rollback migrations
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations rolled back successfully!")
	}
}
