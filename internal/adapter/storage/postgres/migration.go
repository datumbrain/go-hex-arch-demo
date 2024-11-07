package postgres

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migration/*.sql
var sqlFiles embed.FS

const migrationVersion = 1

func (c *Client) MigrateUp() error {
	log.Println("Starting database migrations...")

	// get source instance
	dirInstance, err := iofs.New(sqlFiles, "migration")
	if err != nil {
		return err
	}
	defer dirInstance.Close()

	// get database instance
	sql, err := c.DB().DB()
	if err != nil {
		return err
	}

	dbInstance, err := postgres.WithInstance(sql, &postgres.Config{})
	if err != nil {
		return err
	}

	// create migration object
	m, err := migrate.NewWithInstance("iofs", dirInstance, "postgres", dbInstance)
	if err != nil {
		return err
	}

	// run migrations
	err = m.Migrate(migrationVersion)
	if err == migrate.ErrNoChange {
		log.Println("No change in database required, migrations skipped...")
		return nil
	} else if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	log.Println("Database migrations completed...")
	return nil
}
