package repo

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/bookmark/crud"
	"github.com/dingu27/bookmark/db"
	"github.com/dingu27/bookmark/model"
)

//GetBookmarks ...
func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, conerr := db.GetMongoCon()
	if conerr != nil {
		json.NewEncoder(w).Encode(conerr)
		return
	}

	obj := crud.NewMongo(db, "bookmark")

	var readmark model.Readmark
	json.NewDecoder(r.Body).Decode(&readmark)

	user, err := obj.FindByEmail(readmark.Email)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
