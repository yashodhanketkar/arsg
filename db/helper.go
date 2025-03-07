package db

import (
	"log"
	"os"
	"path/filepath"
)

func createTables(filename string) {
	path := filepath.Join("db", "schema", filename)
	schemeTasks, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(string(schemeTasks))
	if err != nil {
		log.Fatal(err)
	}
}
