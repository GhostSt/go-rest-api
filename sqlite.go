package main

import (
	"database/sql"
	"errors"

	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase(config *yaml.File) (*sql.DB, error) {
	version, err := config.Get("database.version")

	if err != nil {
		return nil, errors.New("Database version doesn't set in configuration")
	}

	path, err := config.Get("database.path")

	if err != nil {
		return nil, errors.New("Database path doesn't set in configuration")
	}

	db, err := sql.Open(version, path)

	if err != nil {
		panic(err)
	}

	return db, nil
}
