package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	registry = createRegistry()

	loadConfig(registry)

	db, err := SetupDatabase(config)

	fmt.Println("test")
	fmt.Println(db)
	fmt.Println(err)
}

func createRegistry () *registry {
	return &registry{}
}

func loadConfig(registry *registry) {
	file, err := yaml.ReadFile("config.yml")

	if (err != nil) {
		panic(err)
	}

	registry.config = file
}

type registry struct {
	config *yaml.File
}