package main

import (
	"github.com/yashodhanketkar/arsg/src/api"
	"github.com/yashodhanketkar/arsg/src/db"
	// "github.com/yashodhanketkar/arsg/src/ui"
)

func main() {
	DB := db.InitDB()
	defer DB.Close()
	// ui.TeaUI(DB)
	api.Serve(DB)
}
