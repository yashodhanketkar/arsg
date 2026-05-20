package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yashodhanketkar/arsg/src/db"
	"github.com/yashodhanketkar/arsg/src/util"
)

func handleMocker(t *testing.T) *Server {
	t.Helper()

	memDB := util.MockDB(t)

	animeRating := db.Rating{
		Name:     "mock anime rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock rating value",
	}

	mangaRating := db.Rating{
		Name:     "mock manga rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock rating value",
	}

	lnRating := []db.Rating{
		{
			Name:     "mock ln rating",
			Art:      8.5,
			Support:  8.5,
			Plot:     8.5,
			Bias:     8.5,
			Rating:   "8.5",
			Comments: "8.5 mock rating value",
		},
		{
			Name:     "mock ln rating 2",
			Art:      7.5,
			Support:  7.5,
			Plot:     7.5,
			Bias:     7.5,
			Rating:   "7.5",
			Comments: "7.5 mock rating 2 value",
		},
	}

	db.AddRatings(memDB, animeRating, "anime")
	db.AddRatings(memDB, mangaRating, "manga")
	for _, r := range lnRating {
		db.AddRatings(memDB, r, "lightnovel")
	}

	return NewServer(memDB)
}

// func handleMocker(t *testing.T) *Server {
// 	t.Helper()
//
// 	memDB := util.MockDB(t)
//
// 	origDB := db.ConnectDB
// 	t.Cleanup(func() {
// 		db.ConnectDB = origDB
// 	})
//
// 	db.ConnectDB = func() *sql.DB {
// 		animeRating := db.Rating{
// 			Name:     "mock anime rating",
// 			Art:      8.5,
// 			Support:  8.5,
// 			Plot:     8.5,
// 			Bias:     8.5,
// 			Rating:   "8.5",
// 			Comments: "8.5 mock rating value",
// 		}
//
// 		mangaRating := db.Rating{
// 			Name:     "mock manga rating",
// 			Art:      8.5,
// 			Support:  8.5,
// 			Plot:     8.5,
// 			Bias:     8.5,
// 			Rating:   "8.5",
// 			Comments: "8.5 mock rating value",
// 		}
//
// 		lnRating := []db.Rating{
// 			{
// 				Name:     "mock ln rating",
// 				Art:      8.5,
// 				Support:  8.5,
// 				Plot:     8.5,
// 				Bias:     8.5,
// 				Rating:   "8.5",
// 				Comments: "8.5 mock rating value",
// 			},
// 			{
// 				Name:     "mock ln rating 2",
// 				Art:      7.5,
// 				Support:  7.5,
// 				Plot:     7.5,
// 				Bias:     7.5,
// 				Rating:   "7.5",
// 				Comments: "7.5 mock rating 2 value",
// 			},
// 		}
//
// 		db.AddRatings(memDB, animeRating, "anime")
// 		db.AddRatings(memDB, mangaRating, "manga")
// 		for _, r := range lnRating {
// 			db.AddRatings(memDB, r, "lightnovel")
// 		}
//
// 		return memDB
// 	}
//
// 	return NewServer(memDB)
// }

func TestCalculator(t *testing.T) {
	// handleMocker(t)
	s := handleMocker(t)

	t.Run("should result in 200 and correct output", func(t *testing.T) {
		tests := []struct {
			testingFor           string
			pathValue            string
			expectedStatus       int
			expectedContentCount int
		}{
			{
				testingFor:           "anime",
				pathValue:            "anime",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 1,
			},
			{
				testingFor:           "manga",
				pathValue:            "manga",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 1,
			},
			{
				testingFor:           "light novels",
				pathValue:            "lightnovel",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 2,
			},
		}

		for _, tt := range tests {
			req := httptest.NewRequest(http.MethodGet, "/list/"+tt.pathValue, nil)
			w := httptest.NewRecorder()
			req.SetPathValue("content_type", tt.pathValue)

			s.listHandler(w, req)

			resp := w.Result()
			resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf(
					"[%s] Expected status code %d, recieved %d",
					tt.testingFor,
					tt.expectedStatus,
					resp.StatusCode,
				)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("[%s] Error reading body: %s", tt.testingFor, err)
			}

			var ratings []db.Rating
			if err := json.Unmarshal(bodyBytes, &ratings); err != nil {
				t.Errorf("[%s] Error unmarshalling body: %s", tt.testingFor, err)
			}

			if len(ratings) != tt.expectedContentCount {
				t.Errorf(
					"[%s] Expected content count %d, recieved %d",
					tt.testingFor,
					tt.expectedContentCount,
					len(ratings),
				)
			}
		}
	})

	t.Run("should result in 400", func(t *testing.T) {
		tests := []struct {
			testingFor           string
			pathValue            string
			expectedStatus       int
			expectedContentCount int
		}{
			{
				testingFor:           "movies",
				pathValue:            "movies",
				expectedStatus:       http.StatusBadRequest,
				expectedContentCount: 0,
			},
			{
				testingFor:           "drama",
				pathValue:            "drama",
				expectedStatus:       http.StatusBadRequest,
				expectedContentCount: 0,
			},
		}

		for _, tt := range tests {
			req := httptest.NewRequest(http.MethodGet, "/list/"+tt.pathValue, nil)
			w := httptest.NewRecorder()
			req.SetPathValue("content_type", tt.pathValue)

			s.listHandler(w, req)

			resp := w.Result()
			resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf(
					"[%s] Expected status code %d, recieved %d",
					tt.testingFor,
					tt.expectedStatus,
					resp.StatusCode,
				)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("[%s] Error reading body: %s", tt.testingFor, err)
			}

			if !strings.Contains(string(bodyBytes), "Invalid content type: "+tt.pathValue) {
				t.Errorf(
					"[%s] Expected error message, recieved %s",
					tt.testingFor,
					string(bodyBytes),
				)
			}
		}
	})
}
