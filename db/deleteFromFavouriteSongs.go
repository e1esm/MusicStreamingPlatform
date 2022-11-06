package db

import (
	"database/sql"
)

func DeleteFromFavourites(db *sql.DB, songTitle string, userName string) error {
	_, err := db.Exec("DELETE FROM favouritesongs WHERE song_id = (SELECT song_id FROM songs WHERE song_name = $1) AND user_id = (SELECT user_id FROM users WHERE username = $2);", songTitle, userName)
	return err
}
