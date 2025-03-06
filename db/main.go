package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Rating struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Art      float32 `json:"art"`
	Support  float32 `json:"support"`
	Plot     float32 `json:"plot"`
	Bias     float32 `json:"bias"`
	Rating   float32 `json:"rating"`
	Comments string  `json:"comments"`
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "args.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitiDB() {
	var err error
	if DB, err = sql.Open("sqlite3", "args.db"); err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	createTables("args.sql")
}

func AddRatings(db *sql.DB, ratings Rating) {
	stmt, err := db.Prepare(
		"INSERT INTO rating (name, art, support, plot, bias, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ratings.Name,
		ratings.Art,
		ratings.Support,
		ratings.Plot,
		ratings.Bias,
		ratings.Rating,
		ratings.Comments,
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}

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
