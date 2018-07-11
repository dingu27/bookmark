package repo

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/bookmark/crud"
	"github.com/dingu27/bookmark/db"
	"github.com/dingu27/bookmark/model"
)

//CheckBookmark ...
func CheckBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, conerr := db.GetMongoCon()
	if conerr != nil {
		json.NewEncoder(w).Encode(conerr)
	}

	obj := crud.NewMongo(db, "bookmark")

	var readmark model.Readmark
	json.NewDecoder(r.Body).Decode(&readmark)

	_, finderr := obj.FindByEmailURL(&readmark)
	if finderr != nil {
		json.NewEncoder(w).Encode("no")
		return
	}
	json.NewEncoder(w).Encode("yes")
}
