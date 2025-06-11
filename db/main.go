package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB

	basePath = filepath.Join(os.Getenv("HOME"), ".local/share/args/lib")
	dbPath   = filepath.Join(basePath, "arsg.db")
)

type Rating struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Art      float32 `json:"art"`
	Support  float32 `json:"support"`
	Plot     float32 `json:"plot"`
	Bias     float32 `json:"bias"`
	Rating   string  `json:"rating"`
	Comments string  `json:"comments"`
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB() {
	var err error

	if tempErr := os.Mkdir(filepath.Dir(basePath), 0755); !os.IsExist(tempErr) {
		log.Fatal(tempErr)
	}

	if DB, err = sql.Open("sqlite3", dbPath); err != nil {
		log.Fatal(err)
	}

	defer DB.Close()
	createTables(DB, filepath.Join(basePath, "schema/arsg.sql"))
}

func AddRatings(db *sql.DB, ratings Rating) {
	rows, err := db.Query(
		"INSERT INTO rating (name, art, support, plot, bias, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)",
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
	defer rows.Close()
	rows.Next()
}

func ListRatings(db *sql.DB) []Rating {
	ratings := make([]Rating, 0)
	rows, err := db.Query("SELECT * FROM rating ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var rating Rating
		err = rows.Scan(
			&rating.ID,
			&rating.Name,
			&rating.Art,
			&rating.Support,
			&rating.Plot,
			&rating.Bias,
			&rating.Rating,
			&rating.Comments,
		)
		if err != nil {
			log.Fatal(err)
		}
		ratings = append(ratings, rating)
	}
	return ratings
}

func createTables(db *sql.DB, path string) {
	schemeTasks, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schemeTasks))
	if err != nil {
		log.Fatal(err)
	}
}

func ExportData(db *sql.DB) {
	data := ListRatings(db)
	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("export.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
