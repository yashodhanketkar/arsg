package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

const schema = `
CREATE TABLE IF NOT EXISTS rating (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	art REAL,
	support REAL,
	plot REAL,
	bias REAL,
	rating TEXT,
	comments TEXT
);
`

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	schemaPath := filepath.Join(t.TempDir(), "schema.sql")
	err = os.WriteFile(schemaPath, []byte(schema), 0644)

	if err != nil {
		t.Fatalf("failed to write schema file: %v", err)
	}
	createTables(db, schemaPath)

	return db
}

func TestAddListRatings(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	rating := Rating{
		Name:     "mock rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock rating value",
	}

	AddRatings(db, rating)
	result := ListRatings(db)

	if n := len(result); n != 1 {
		t.Fatalf("want 1, got %d", n)
	}

	got := result[0]
	if got.Name != rating.Name {
		t.Errorf("want %q, got %q", rating.Name, got.Name)
	}

	if got.Art != rating.Art {
		t.Errorf("want %f, got %f", rating.Art, got.Art)
	}

	if got.Support != rating.Support {
		t.Errorf("want %f, got %f", rating.Support, got.Support)
	}

	if got.Plot != rating.Plot {
		t.Errorf("want %f, got %f", rating.Plot, got.Plot)
	}

	if got.Bias != rating.Bias {
		t.Errorf("want %f, got %f", rating.Bias, got.Bias)
	}

	if got.Rating != rating.Rating {
		t.Errorf("want %q, got %q", rating.Rating, got.Rating)
	}

	if got.Comments != rating.Comments {
		t.Errorf("want %q, got %q", rating.Comments, got.Comments)
	}
}
