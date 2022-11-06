package db

import (
	"database/sql"
	"encoding/json"
	"log"
)

func AddNewFavouriteTrack(sql *sql.DB, songTitle string, currentUser string) ([]byte, error) {

	log.Println(songTitle, currentUser)

	_, err := sql.Exec("INSERT INTO favouritesongs (song_id, user_id) VALUES ((SELECT song_id FROM songs WHERE song_name = $1), (SELECT user_id FROM users where username = $2));", songTitle, currentUser)

	if err != nil {
		log.Fatal(err)
	}

	return json.Marshal(songTitle)
}
