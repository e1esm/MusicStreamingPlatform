package handler

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	db2 "musicStreamingPlatform/db"
	"net/http"
)

func GetQueriedSongs(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songTitle := vars["songTitle"]
	//songTitle := r.URL.Query().Get("songTitle")
	jsonRepr, err := db2.GetAllSongs(db, songTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(jsonRepr)
	w.Write(jsonRepr)

}
