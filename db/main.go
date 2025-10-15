package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

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

func addRating(tx *sql.Tx, ratings Rating) (int64, error) {
	stmt, err := tx.Prepare(
		"INSERT INTO rating (name, art, support, plot, bias, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		ratings.Name,
		ratings.Art,
		ratings.Support,
		ratings.Plot,
		ratings.Bias,
		ratings.Rating,
		ratings.Comments,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func addToIndex(tx *sql.Tx, id int64, contentType string) error {
	_, err := tx.Exec("INSERT INTO "+contentType+" (rating_id) VALUES (?)", id)
	return err
}

func AddRatings(db *sql.DB, ratings Rating, contentType string) error {
	tx, err := db.Begin()
	defer tx.Rollback()

	id, err := addRating(tx, ratings)
	if err != nil {
		return err
	}
	addToIndex(tx, id, contentType)

	return tx.Commit()
}

func ListRatings(db *sql.DB, contentType string) []Rating {
	ratings := make([]Rating, 0)
	querySQL := fmt.Sprint(
		"SELECT a.* FROM ",
		contentType,
		" AS c LEFT JOIN rating AS a ON a.id = c.rating_id ORDER BY a.name",
	)

	rows, err := db.Query(querySQL)
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

func ExportData(db *sql.DB, exportPath string) {
	var data JData

	data.Anime = ListRatings(db, "anime")
	data.Manga = ListRatings(db, "manga")
	data.LightNovel = ListRatings(db, "lightnovel")

	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	if err = os.MkdirAll(filepath.Dir(exportPath), 0755); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(exportPath)
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
