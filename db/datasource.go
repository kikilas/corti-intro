package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"log"
	"strings"
)

type Config struct {
	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
	DbArgs string
}

func (c Config) ConnectionString() (connection string) {
	connection = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName, c.DbArgs)

	return
}

func Connect(cfg Config) (*sqlx.DB, error) {
	connectionString := cfg.ConnectionString()
	log.Printf("Connecting to %s", connectionString)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	// Create a new mapper which will use the struct field tag "json" instead of "db"
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	return db, err
}
