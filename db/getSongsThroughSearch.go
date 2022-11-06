package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"musicStreamingPlatform/model"
)

func GetAllSongs(db *sql.DB, title string) ([]byte, error) {
	query := `SELECT
  song.song_id, song.album_id, 
  song.song_name, album.album_cover_image, album.author, album.album_title
FROM
  songs AS song
  INNER JOIN albums AS album
    ON album.album_id = song.song_id
WHERE similarity(song.song_name, $1) > 0.4;`

	row, err := db.Query(query, title)
	if err != nil {
		log.Fatal(err)
	}

	var retrievedSongs []model.Song

	for row.Next() {
		var songId int
		var title string
		var albumId int
		var coverageImage string
		var artist string
		var albumTitle string

		err := row.Scan(&songId, &albumId, &title, &coverageImage, &artist, &albumTitle)
		if err != nil {
			log.Fatal(err)
		}
		retrievedSongs = append(retrievedSongs, model.Song{SongID: songId, SongTitle: title, CurrentAlbum: model.Album{AlbumID: albumId, AlbumTitle: albumTitle, AlbumCoverImage: coverageImage, Author: artist}})
	}
	return json.Marshal(retrievedSongs)
}
