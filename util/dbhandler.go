package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func MockDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	schemaPath := filepath.Join(t.TempDir(), "schema.sql")
	schemaSQL := fmt.Sprintf(MockDBShema, "rating")
	err = os.WriteFile(schemaPath, []byte(schemaSQL), 0644)

	if err != nil {
		t.Fatalf("failed to write schema file: %v", err)
	}
	CreateTables(db, schemaPath)

	t.Cleanup(func() {
		_, _ = db.Exec("DELETE FROM anime;")
		_, _ = db.Exec("DELETE FROM manga;")
		_, _ = db.Exec("DELETE FROM rating;")

		db.Close()
	})

	return db
}

func CreateTables(db *sql.DB, path string) {
	schemeTasks, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schemeTasks))
	if err != nil {
		log.Fatal(err)
	}
}
