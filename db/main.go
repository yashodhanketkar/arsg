package db

import (
	"database/sql"
	"log"

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
	Rating   string  `json:"rating"`
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
