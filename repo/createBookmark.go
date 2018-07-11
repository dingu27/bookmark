package repo

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/bookmark/crud"
	"github.com/dingu27/bookmark/db"
	"github.com/dingu27/bookmark/model"
)

//CreateBookmark ...
func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, conerr := db.GetMongoCon()

	if conerr != nil {
		json.NewEncoder(w).Encode(conerr)
	}

	obj := crud.NewMongo(db, "bookmark")

	var readmark model.Readmark
	json.NewDecoder(r.Body).Decode(&readmark)

	if readmark.Bookmark == "no" {
		saverr := obj.Save(&readmark)
		if saverr != nil {
			json.NewEncoder(w).Encode(saverr)
			return
		}
		json.NewEncoder(w).Encode("Success Bookmarked")
		return
	}
	user, finderr := obj.FindByEmailURL(&readmark)
	if finderr != nil {
		json.NewEncoder(w).Encode("finderr")
		return
	}
	delerr := obj.Delete(user)
	if delerr != nil {
		json.NewEncoder(w).Encode("delerr")
		return
	}
	json.NewEncoder(w).Encode("Success UnBookmarked")
}
