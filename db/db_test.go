package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type ratingCtx struct {
	id       int64
	ratingId int64
}

const schema = `
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	art REAL,
	support REAL,
	plot REAL,
	bias REAL,
	rating TEXT,
	comments TEXT
);

CREATE TABLE IF NOT EXISTS anime (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  rating_id integer NOT NULL,
  FOREIGN KEY (rating_id) REFERENCES rating (id)
);

CREATE TABLE IF NOT EXISTS manga (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  rating_id integer NOT NULL,
  FOREIGN KEY (rating_id) REFERENCES rating (id)
);
`

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	schemaPath := filepath.Join(t.TempDir(), "schema.sql")
	schemaSQL := fmt.Sprintf(schema, "rating")
	err = os.WriteFile(schemaPath, []byte(schemaSQL), 0644)

	if err != nil {
		t.Fatalf("failed to write schema file: %v", err)
	}
	createTables(db, schemaPath)

	return db
}

func TestAddListRatings(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	animeRating := Rating{
		Name:     "mock anime rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock rating value",
	}

	mangaRating := Rating{
		Name:     "mock manga rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock manga value",
	}

	AddRatings(db, animeRating, "anime")
	testAddRating(t, db, animeRating, "anime")
	testAnimeRating(t, db, ratingCtx{id: 1, ratingId: 1})

	AddRatings(db, mangaRating, "manga")
	testAddRating(t, db, mangaRating, "manga")
	testMangaRating(t, db, ratingCtx{id: 1, ratingId: 2})
}

func testAddRating(t *testing.T, db *sql.DB, rating Rating, contentType string) {
	result := ListRatings(db, contentType)

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

func testAnimeRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	var id, ratingID int64

	rows, err := db.Query("SELECT * FROM anime")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	rows.Next()

	rows.Scan(&id, &ratingID)

	if id != ctx.id {
		t.Errorf("want 1, got %d", id)
	}

	if ratingID != ctx.ratingId {
		t.Errorf("want 1, got %d", ratingID)
	}
}

func testMangaRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	var id, ratingID int64

	rows, err := db.Query("SELECT * FROM manga")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	rows.Next()

	rows.Scan(&id, &ratingID)

	if id != ctx.id {
		t.Errorf("want 1, got %d", id)
	}

	if ratingID != ctx.ratingId {
		t.Errorf("want 1, got %d", ratingID)
	}
}
