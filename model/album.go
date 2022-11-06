package model

type Album struct {
	AlbumID         int    `json:"albumId"`
	AlbumTitle      string `json:"AlbumTitle"`
	AlbumCoverImage string `json:"album_cover_image"`
	SongsInAlbum    []Song `json:"songs_in_album"`
	Author          string `json:"album_author"`
}
