package db

import (
	"database/sql"
	"log"
)

func SetUpTables(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Albums (album_id SERIAL NOT NULL PRIMARY KEY, album_cover_image TEXT, Author TEXT, album_title TEXT UNIQUE);")
	if err != nil {
		log.Fatal(err.Error() + "Setting up database")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Songs (album_id INT, song_id SERIAl NOT NULL PRIMARY KEY, song_name TEXT, CONSTRAINT fk_album FOREIGN KEY(album_id) REFERENCES Albums(album_id));")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Users (user_id SERIAl NOT NULL PRIMARY KEY, Username TEXT);")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS FavouriteSongs(song_id INT, user_id INT, CONSTRAINT fk_liked_song_id FOREIGN KEY(song_id) REFERENCES Songs(song_id), CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES Users(user_id));")
	if err != nil {
		log.Fatal(err.Error())
	}

}
