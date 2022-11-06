CREATE TABLE IF NOT EXISTS Albums (album_id SERIAL NOT NULL PRIMARY KEY, album_cover_image TEXT, Author TEXT, album_title TEXT UNIQUE);

CREATE TABLE IF NOT EXISTS Songs (album_id INT, song_id SERIAl NOT NULL PRIMARY KEY, song_name TEXT, CONSTRAINT fk_album FOREIGN KEY(album_id) REFERENCES Albums(album_id));

CREATE TABLE IF NOT EXISTS Users (user_id SERIAl NOT NULL PRIMARY KEY, Username TEXT);

CREATE TABLE IF NOT EXISTS FavouriteSongs(song_id INT, user_id INT, CONSTRAINT fk_liked_song_id FOREIGN KEY(song_id) REFERENCES Songs(song_id), CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES Users(user_id));