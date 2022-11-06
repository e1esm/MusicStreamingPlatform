package model

type User struct {
	UserID    int    `json:"userID"`
	Username  string `json:"username"`
	likedSong Song   `json:"likedSong"`
	Role      string `json:"Role"`
}
