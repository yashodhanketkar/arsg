package api

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yashodhanketkar/arsg/src/db"
)

func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(autoCleanBody)

	// inject database connection
	DB := db.ConnectDB()
	defer DB.Close()

	// main router
	ratingRouter(r, DB)

	// start server
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}

func autoCleanBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer func() {
				io.Copy(io.Discard, r.Body)
				r.Body.Close()
			}()

			next.ServeHTTP(w, r)
		}
	})
}
