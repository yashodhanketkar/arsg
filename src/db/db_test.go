package db

import (
	"database/sql"
	"testing"

	"github.com/yashodhanketkar/arsg/src/util"
)

type ratingCtx struct {
	id       int64
	ratingId int64
}

const (
	wantOneGotMany = util.WantOneGotMany
	wantQGotQ      = util.WantQGotQ
	wantFGotF      = util.WantFGotF
)

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

	AddRatings(db, animeRating, "anime")
	testAddRating(t, db, animeRating, "anime")
	testAnimeRating(t, db, ratingCtx{id: 1, ratingId: 1})

	AddRatings(db, mangaRating, "manga")
	testAddRating(t, db, mangaRating, "manga")
	testMangaRating(t, db, ratingCtx{id: 1, ratingId: 2})
}

func testAddRating(t *testing.T, db *sql.DB, rating Rating, contentType string) {
	result := ListRatings(db, contentType)

	if n := len(result); n != 1 {
		t.Fatalf(wantOneGotMany, n)
	}

	got := result[0]
	if got.Name != rating.Name {
		t.Errorf(wantQGotQ, rating.Name, got.Name)
	}

	if got.Art != rating.Art {
		t.Errorf(wantFGotF, rating.Art, got.Art)
	}

	if got.Support != rating.Support {
		t.Errorf(wantFGotF, rating.Support, got.Support)
	}

	if got.Plot != rating.Plot {
		t.Errorf(wantFGotF, rating.Plot, got.Plot)
	}

	if got.Bias != rating.Bias {
		t.Errorf(wantFGotF, rating.Bias, got.Bias)
	}

	if got.Rating != rating.Rating {
		t.Errorf(wantQGotQ, rating.Rating, got.Rating)
	}

	if got.Comments != rating.Comments {
		t.Errorf(wantQGotQ, rating.Comments, got.Comments)
	}
}

func testAnimeRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	var id, ratingID int64

	rows, err := db.Query("SELECT * FROM anime")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	rows.Next()

	rows.Scan(&id, &ratingID)

	if id != ctx.id {
		t.Errorf(wantOneGotMany, id)
	}

	if ratingID != ctx.ratingId {
		t.Errorf(wantOneGotMany, ratingID)
	}
}

func testMangaRating(t *testing.T, db *sql.DB, ctx ratingCtx) {
	var id, ratingID int64

	rows, err := db.Query("SELECT * FROM manga")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	rows.Next()

	rows.Scan(&id, &ratingID)

	if id != ctx.id {
		t.Errorf(wantOneGotMany, id)
	}

	if ratingID != ctx.ratingId {
		t.Errorf(wantOneGotMany, ratingID)
	}
}
