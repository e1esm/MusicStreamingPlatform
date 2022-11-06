package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"musicStreamingPlatform/model"
)

func QueryForFavouriteSongs(db *sql.DB, username string) ([]byte, error) {
	log.Println(username)
	row, err := db.Query("SELECT song.song_id, song.album_id,  song.song_name, album.album_cover_image, album.author, album.album_title FROM  songs AS song  INNER JOIN albums AS album ON album.album_id = song.song_id INNER JOIN favouritesongs f on song.song_id = f.song_id AND f.user_id = (SELECT user_id FROM users WHERE username = $1);", username)

	if err != nil {
		log.Fatal(err.Error())
	}

	var favouriteSongs []model.Song

	for row.Next() {
		var songId int
		var albumId int
		var songName string
		var albumCoverage string
		var author string
		var albumTitle string
		err := row.Scan(&songId, &albumId, &songName, &albumCoverage, &author, &albumTitle)
		if err != nil {
			return nil, err
		}

		log.Printf("SongID: %d, AlbumID: %d, SongName: %s, Author: %s", songId, albumId, songName, author)
		favouriteSongs = append(favouriteSongs, model.Song{SongID: songId, SongTitle: songName, CurrentAlbum: model.Album{AlbumID: albumId, AlbumTitle: albumTitle, AlbumCoverImage: albumCoverage, Author: author}})
	}
	return json.Marshal(favouriteSongs)
}
