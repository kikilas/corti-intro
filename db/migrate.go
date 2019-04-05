package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq" // import postgres driver
)

func Migrate(db *sql.DB) error {
	driver, err := cockroachdb.WithInstance(db, &cockroachdb.Config{})
	if err != nil {
		log.Fatalf("An error occurred while connecting to DB %v", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("Migration failed %v", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
		return err
	}

	version, _, _ := m.Version()
	log.Printf("Database schema version: %d", version)

	return nil
}
