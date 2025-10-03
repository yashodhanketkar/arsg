package db

import (
	"database/sql"
	"os"
	"path/filepath"
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

type JData struct {
	Anime      []Rating `json:"anime"`
	Manga      []Rating `json:"manga"`
	LightNovel []Rating `json:"lightnovel"`
}
