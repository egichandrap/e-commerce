package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warningln(err)
	}

	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Panicln(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panicln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+os.Getenv("MIGRATION_DIR"),
		"postgres",
		driver,
	)
	if err != nil {
		log.Panicln(err)
	}

	err = m.Up()
	if err != nil {
		log.Warningln(err)
	}

	log.Infoln("success migrate")
	os.Exit(0)
}
