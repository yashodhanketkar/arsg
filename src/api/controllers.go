package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	db *sql.DB
}

// Create a new server instance with the given database connection
func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func ratingRouter(r *chi.Mux, db *sql.DB) {
	s := NewServer(db)

	r.Get("/list/{content_type}", s.listHandler)
	r.Post("/calc", s.calcHandler)
	r.Post("/add/{content_type}", s.addHandler)
}

// returns list of ratings for a given content type.
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	res, err := contentList(r.PathValue("content_type"), s.db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// returns calculated rating for a given content type without saving it to the
// database.
// WARN:temporary endpoint to show rating in the UI.
// Will be polled higher hence highly taxing but will remove once I have
// efficient way to show current calculated rating.
func (s *Server) calcHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	res, err := calculateResult(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// calculate and stores rating for a given content type in the database.
func (s *Server) addHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	err := addRating(r.PathValue("content_type"), body, s.db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Success"))
}
