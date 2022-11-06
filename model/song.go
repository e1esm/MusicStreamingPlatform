package model

type Song struct {
	SongID       int    `json:"songID"`
	SongTitle    string `json:"songTitle"`
	CurrentAlbum Album  `json:"Album"`
	ArtistName   string `json:"Artist"`
}
