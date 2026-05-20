package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CalcResponse struct {
	Title    string  `json:"title"`
	Comments string  `json:"comments"`
	Art      float32 `json:"art"`
	Cast     float32 `json:"cast"`
	Plot     float32 `json:"plot"`
	Bias     float32 `json:"bias"`
	Rating   float32 `json:"rating"`
}

func Serve(DB *sql.DB) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(autoCleanBody)

	// inject database connection
	defer DB.Close()

	// main router
	ratingRouter(r, DB)

	// start server
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
