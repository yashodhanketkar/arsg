package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/yashodhanketkar/arsg/src/db"
	"github.com/yashodhanketkar/arsg/src/util"
)

func handleMocker(t *testing.T) *chi.Mux {
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

	s := NewServer(memDB)
	r := chi.NewRouter()
	ratingRouter(r, s.db)

	return r
}

func TestRatingRouter(t *testing.T) {
	r := handleMocker(t)
	tests := []struct {
		name           string
		method         string
		targetUrl      string
		body           interface{}
		expectedStatus int
	}{
		{
			name:           "GET list maps correctly",
			method:         http.MethodGet,
			targetUrl:      "/list/anime",
			body:           nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST calc maps correctly",
			method:         http.MethodPost,
			targetUrl:      "/calc",
			body:           map[string]string{"Title": "test"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:      "POST add maps correctly",
			method:    http.MethodPost,
			targetUrl: "/add/anime",
			body: map[string]string{
				"Title":    "test",
				"Comments": "test",
				"Art":      "10.0",
				"Cast":     "10.0",
				"Plot":     "10.0",
				"Bias":     "10.0",
			},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := testPayloadParser(t, tt.body)

			req := httptest.NewRequest(tt.method, tt.targetUrl, &payload)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestListHandler(t *testing.T) {
	r := handleMocker(t)

	t.Run("should result in 200 and correct output", func(t *testing.T) {
		tests := []struct {
			name                 string
			pathValue            string
			expectedStatus       int
			expectedContentCount int
		}{
			{
				name:                 "anime",
				pathValue:            "anime",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 1,
			},
			{
				name:                 "manga",
				pathValue:            "manga",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 1,
			},
			{
				name:                 "light novels",
				pathValue:            "lightnovel",
				expectedStatus:       http.StatusOK,
				expectedContentCount: 2,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/list/"+tt.pathValue, nil)
				req.SetPathValue("content_type", tt.pathValue)

				w := httptest.NewRecorder()

				r.ServeHTTP(w, req)
				assert.Equal(t, tt.expectedStatus, w.Code)

				bodyBytes, err := io.ReadAll(w.Body)
				if assert.NoError(t, err, "Error reading body: %s", err) {

					var ratings []db.Rating
					err = json.Unmarshal(bodyBytes, &ratings)

					if assert.NoError(t, err, "Error unmarshalling body: %s", err) {
						assert.Len(t, ratings, tt.expectedContentCount)
					}
				}
			})
		}
	})

	t.Run("should result in 400", func(t *testing.T) {
		tests := []struct {
			name                 string
			pathValue            string
			expectedStatus       int
			expectedContentCount int
			expectedErr          string
		}{
			{
				name:           "should fail with incorrect content type - movies",
				pathValue:      "movies",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Invalid content type: movies",
			},
			{
				name:           "should fail with incorrect content type - drama",
				pathValue:      "drama",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Invalid content type: drama",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/list/"+tt.pathValue, nil)
				req.SetPathValue("content_type", tt.pathValue)

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)
				testBodyError(t, w, tt.expectedErr)
			})
		}
	})
}

func TestCalcHadler(t *testing.T) {
	r := handleMocker(t)

	t.Run("should result in 200 and correct output", func(t *testing.T) {
		tests := []struct {
			name           string
			data           interface{}
			expectedStatus int
			expectedScore  float32
		}{
			{
				name: "calculate for full score",
				data: map[string]string{
					"Title":    "test",
					"Comments": "test",
					"Art":      "10.0",
					"Cast":     "10.0",
					"Plot":     "10.0",
					"Bias":     "10.0",
				},
				expectedStatus: http.StatusOK,
				expectedScore:  10.0,
			},
			{
				name: "calculate for general score",
				data: map[string]string{
					"Title":    "test 2",
					"Comments": "test 2",
					"Art":      "8.0",
					"Cast":     "9.0",
					"Plot":     "8.0",
					"Bias":     "7.0",
				},
				expectedStatus: http.StatusOK,
				expectedScore:  8.2,
			},
			{
				name: "calculate for low score",
				data: map[string]string{
					"Title":    "test",
					"Comments": "test",
					"Art":      "0.1",
					"Cast":     "0.1",
					"Plot":     "0.1",
					"Bias":     "0.1",
				},
				expectedStatus: http.StatusOK,
				expectedScore:  0.1,
			},
		}

		for _, tt := range tests {
			payload := testPayloadParser(t, tt.data)
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPost, "/calc", &payload)

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)

				var rating CalcResponse
				err := json.NewDecoder(w.Body).Decode(&rating)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expectedScore, rating.Rating)
				}
			})
		}
	})

	t.Run("should result in 400", func(t *testing.T) {
		tests := []struct {
			name           string
			data           interface{}
			expectedStatus int
			expectedErr    string
		}{
			{
				name: "calculate for zero score",
				data: map[string]string{
					"Title":    "test",
					"Comments": "test",
					"Art":      "0",
					"Cast":     "0",
					"Plot":     "0",
					"Bias":     "0",
				},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "All zero values provided",
			},
			{
				name:           "should fail with missing title",
				data:           map[string]string{},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Missing title",
			},
			{
				name:           "should fail with missing art",
				data:           map[string]string{"Title": "test"},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Missing parameter - Art",
			},
			{
				name:           "should fail with missing cast",
				data:           map[string]string{"Title": "test", "Art": "1"},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Missing parameter - Cast",
			},
			{
				name:           "should fail with missing plot",
				data:           map[string]string{"Title": "test", "Art": "1", "Cast": "1"},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Missing parameter - Plot",
			},
			{
				name: "should fail with missing bias",
				data: map[string]string{
					"Title": "test",
					"Art":   "1",
					"Cast":  "1",
					"Plot":  "1",
				},
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Missing parameter - Bias",
			},
		}

		for _, tt := range tests {
			payload := testPayloadParser(t, tt.data)

			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPost, "/calc", &payload)

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)
				testBodyError(t, w, tt.expectedErr)
			})
		}
	})
}

func TestAddHandler(t *testing.T) {
	r := handleMocker(t)

	t.Run("should result in 201 and correct output", func(t *testing.T) {
		tests := []struct {
			name           string
			data           map[string]string
			pathValue      string
			expectedResult db.Rating
		}{
			{
				name: "checking for full score for anime",
				data: map[string]string{
					"Title":    "test",
					"Comments": "test",
					"Art":      "10.0",
					"Cast":     "10.0",
					"Plot":     "10.0",
					"Bias":     "10.0",
				},
				pathValue: "anime",
				expectedResult: db.Rating{
					Name:     "test",
					Art:      10.0,
					Support:  10.0,
					Plot:     10.0,
					Bias:     10.0,
					Rating:   "10.0",
					Comments: "test",
				},
			},
			{
				name: "checking for general score for manga",
				data: map[string]string{
					"Title":    "test-manga",
					"Comments": "test-manga-comments",
					"Art":      "7.0",
					"Cast":     "7.0",
					"Plot":     "7.0",
					"Bias":     "7.0",
				},
				pathValue: "manga",
				expectedResult: db.Rating{
					Name:     "test-manga",
					Art:      7.0,
					Support:  7.0,
					Plot:     7.0,
					Bias:     7.0,
					Rating:   "7.0",
					Comments: "test-manga-comments",
				},
			},
			{
				name: "checking for low score for light novel",
				data: map[string]string{
					"Title":    "test-lightnovel",
					"Comments": "test-lightnovel-comments",
					"Art":      "0.1",
					"Cast":     "0.1",
					"Plot":     "0.1",
					"Bias":     "0.1",
				},
				pathValue: "lightnovel",
				expectedResult: db.Rating{
					Name:     "test-lightnovel",
					Art:      0.1,
					Support:  0.1,
					Plot:     0.1,
					Bias:     0.1,
					Rating:   "0.1",
					Comments: "test-lightnovel-comments",
				},
			},
		}

		for _, tt := range tests {
			payload := testPayloadParser(t, tt.data)

			req := httptest.NewRequest(http.MethodPost, "/add/"+tt.pathValue, &payload)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusCreated, w.Code)
			var response db.Rating
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
			assert.Equal(t, tt.expectedResult, response)
		}
	})

	t.Run("should result in 400", func(t *testing.T) {
		tests := []struct {
			name        string
			data        map[string]string
			pathValue   string
			expectedErr string
		}{
			{
				name:        "should fail with incorrect content type",
				data:        nil,
				pathValue:   "incorrect",
				expectedErr: "Invalid content type: incorrect",
			},
			{
				name: "should fail with missing art",
				data: map[string]string{
					"Title":    "test",
					"Comments": "test",
					"Cast":     "10.0",
					"Plot":     "10.0",
					"Bias":     "10.0",
				},
				pathValue:   "anime",
				expectedErr: "Missing parameter - Art",
			},
			{
				name:        "should fail with missing title",
				data:        nil,
				pathValue:   "manga",
				expectedErr: `Missing title`,
			},
		}

		for _, tt := range tests {
			payload := testPayloadParser(t, tt.data)

			req := httptest.NewRequest(http.MethodPost, "/add/"+tt.pathValue, &payload)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			testBodyError(t, w, tt.expectedErr)
		}
	})
}

func testBodyError(t *testing.T, w *httptest.ResponseRecorder, expectedErr string) {
	t.Helper()

	bodyBytes, err := io.ReadAll(w.Body)
	if assert.NoError(t, err, "Error reading body: %s", err) {
		got := string(bodyBytes)
		assert.Contains(
			t, got, expectedErr, "Unexpected error message, recieved %s", got,
		)
	}
}

func testPayloadParser(t *testing.T, data any) bytes.Buffer {
	t.Helper()

	var payload bytes.Buffer
	assert.NoError(t, json.NewEncoder(&payload).Encode(data))

	return payload
}
