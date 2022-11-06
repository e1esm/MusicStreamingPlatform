package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	db2 "musicStreamingPlatform/db"
	"musicStreamingPlatform/model"
	"net/http"
)

func ReleaseNewTrack(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	newReleasedTrack := model.Song{}
	json.Unmarshal(body, &newReleasedTrack)

	db2.AddNewTrackToDatabase(db, &newReleasedTrack)

}

func ReleaseNewAlbum(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	newReleasedAlbum := model.Album{}
	json.Unmarshal(body, &newReleasedAlbum)

	db2.AddNewAlbumToDatabase(db, &newReleasedAlbum)
}
