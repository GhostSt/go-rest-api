package main

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/julienschmidt/httprouter"
	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/unrolled/render.v1"
)

type registry struct {
	config *yaml.File
	db     *sql.DB
	render *render.Render
	router *httprouter.Router
}

// Creates router
func createRouter(registry *registry) {
	registry.router = httprouter.New()
}

// Parses configuration file and sets it to Registry
func loadConfig(registry *registry) {
	file, err := yaml.ReadFile("resources/config.yml")

	if (err != nil) {
		panic(err)
	}

	registry.config = file
}

// Sets up application and initialize Registry
func setup() *registry {
	registry := &registry{}

	loadConfig(registry)

	createRender(registry)
	createRouter(registry)

	setupDatabase(registry)

	configureRoutes(registry)

	return registry
}

// Creates render
func createRender(registry *registry)  {
	registry.render = render.New(render.Options{})
}

// Sets up connection to database and sets it to Registry and import initial database schema
func setupDatabase(registry *registry) (error) {
	version, err := registry.config.Get("database.version")

	if err != nil {
		return errors.New("Database version doesn't set in configuration")
	}

	path, err := registry.config.Get("database.path")

	if err != nil {
		return errors.New("Database path doesn't set in configuration")
	}

	db, err := sql.Open(version, path)

	if err != nil {
		return errors.New("Database schema doesn't set in configuration")
	}

	migration_file, err := registry.config.Get("database.schema")

	if err != nil {
		panic(err)
	}

	if _, err = os.Stat(path); os.IsNotExist(err) {
		data, err := ioutil.ReadFile(migration_file)

		if err != nil {
			panic(err)
		}

		_, err = db.Exec(string(data))

		if err != nil {
			panic(err)
		}
	}

	registry.db = db

	return nil
}
