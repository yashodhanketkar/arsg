package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yashodhanketkar/arsg/src/util"
)

type ratingCtx struct {
	id       int64
	ratingId int64
}

func TestAddListRatings(t *testing.T) {
	db := util.MockDB(t)
	defer db.Close()

	animeRating := Rating{
		Name:     "mock anime rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock rating value",
	}

	mangaRating := Rating{
		Name:     "mock manga rating",
		Art:      8.5,
		Support:  8.5,
		Plot:     8.5,
		Bias:     8.5,
		Rating:   "8.5",
		Comments: "8.5 mock manga value",
	}

	t.Run("test anime database insert", func(t *testing.T) {
		AddRatings(db, animeRating, "anime")
		testAddRating(t, db, animeRating, "anime")
		testAnimeRating(t, db, ratingCtx{id: 1, ratingId: 1})
	})

	t.Run("test manga database insert", func(t *testing.T) {
		AddRatings(db, mangaRating, "manga")
		testAddRating(t, db, mangaRating, "manga")
		testMangaRating(t, db, ratingCtx{id: 1, ratingId: 2})
	})
}

func testAddRating(t *testing.T, db *sql.DB, rating Rating, contentType string) {
	result := ListRatings(db, contentType)

	assert.Len(t, result, 1)

	got := result[0]

	assert.Equal(t, rating.Name, got.Name)
	assert.Equal(t, rating.Art, got.Art)
	assert.Equal(t, rating.Support, got.Support)
	assert.Equal(t, rating.Plot, got.Plot)
	assert.Equal(t, rating.Bias, got.Bias)
	assert.Equal(t, rating.Rating, got.Rating)
	assert.Equal(t, rating.Comments, got.Comments)
}

func testAnimeRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	rows, err := db.Query("SELECT * FROM anime")
	if assert.NoError(t, err) {
		defer rows.Close()
	}

	testContentHelper(t, rows, ctx)
}

func testMangaRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	rows, err := db.Query("SELECT * FROM manga")
	if assert.NoError(t, err) {
		defer rows.Close()
	}

	testContentHelper(t, rows, ctx)
}

func testContentHelper(t *testing.T, rows *sql.Rows, ctx ratingCtx) {
	t.Helper()

	var id, ratingID int64

	rows.Next()
	rows.Scan(&id, &ratingID)

	assert.Equal(t, ctx.id, id)
	assert.Equal(t, ctx.ratingId, ratingID)
}
