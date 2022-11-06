package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	db2 "musicStreamingPlatform/db"
	"musicStreamingPlatform/model"
	"net/http"
)

func DeleteFromFavouriteList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	songTitle := r.URL.Query().Get("songTitle")
	username := r.URL.Query().Get("username")
	err := db2.DeleteFromFavourites(db, songTitle, username)

	jsonDeltedElement, _ := json.Marshal(songTitle)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDeltedElement)

}

func AddNewFavouriteInstances(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var song model.Song
	json.Unmarshal(body, &song)

	//songTitle := r.URL.Query().Get("songTitle")
	//username := r.URL.Query().Get("username")

	log.Println(song.SongTitle, song.ArtistName)

	content, err := db2.AddNewFavouriteTrack(db, song.SongTitle, song.ArtistName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//jsonResponse, _ := json.Marshal(fmt.Sprintf("Added: %s", vars["songTitle"]))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

}

func GetAllFavourite(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//username := r.URL.Query().Get("username")

	jsonRepr, err := db2.QueryForFavouriteSongs(db, vars["username"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRepr)

}

/*
func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

*/
