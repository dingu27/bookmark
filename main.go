package main

import (
	"net/http"

	"github.com/dingu27/bookmark/repo"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/bookmark/post", repo.CreateBookmark).Methods("POST")
	router.HandleFunc("/api/bookmark/get", repo.GetBookmarks).Methods("POST")
	router.HandleFunc("/api/bookmark/check", repo.CheckBookmark).Methods("POST")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8000", handler)
}
