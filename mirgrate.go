package databasemigrate

import (
	"errors"
	"net/url"

	"github.com/amacneil/dbmate/pkg/dbmate"
	_ "github.com/amacneil/dbmate/pkg/driver/postgres"
)

func Go(connection string, dump bool) error {
	if connection == "" {
		return errors.New("connection string is missing")
	}

	url, err := url.Parse(connection)
	if err != nil {
		return err
	}

	db := dbmate.New(url)
	db.AutoDumpSchema = dump
	db.MigrationsDir = "./database/migrations"
	db.SchemaFile = "./database/schema.sql"

	return db.CreateAndMigrate()
}
