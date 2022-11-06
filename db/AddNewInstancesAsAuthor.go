package db

import (
	"database/sql"
	"log"
	"musicStreamingPlatform/model"
)

func AddNewAlbumToDatabase(db *sql.DB, album *model.Album) {
	log.Println(album.AlbumTitle, album.AlbumID)

	_, err := db.Exec("INSERT INTO albums (album_cover_image, author, album_title) VALUES ($1, $2, $3)", album.AlbumCoverImage, album.Author, album.AlbumTitle)

	if err != nil {
		log.Fatal(err.Error())
	}
	for i := 0; i < len(album.SongsInAlbum); i++ {
		_, err := db.Exec("INSERT INTO songs (album_id, song_name) VALUES ((SELECT album_id FROM albums WHERE album_title = $1), $2)", album.AlbumTitle, album.SongsInAlbum[i].SongTitle)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func AddNewTrackToDatabase(db *sql.DB, song *model.Song) {
	log.Println(song.SongTitle, song.ArtistName, song.CurrentAlbum)

	_, err := db.Exec("INSERT INTO albums (album_title, author) VALUES ($1, $2)on conflict (album_title, author) do nothing", song.CurrentAlbum.AlbumTitle, song.ArtistName)

	_, err = db.Exec("INSERT INTO songs (album_id, song_name) VALUES ((SELECT album_id FROM albums WHERE album_title = $1),$2);", song.CurrentAlbum.AlbumTitle, song.SongTitle)
	if err != nil {
		log.Fatal(err)
	}

}
