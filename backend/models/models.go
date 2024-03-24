package models

import "time"

// LightNovel represents the LightNovels table.
type LightNovel struct {
	LightNovelID    int      `json:"lightNovelId"`
	Title           string   `json:"title"`
	NumberOfVolumes int      `json:"numberOfVolumes"`
	CoverArtURL     string   `json:"coverArtUrl"`
	Volumes         []Volume `json:"volumes"`
	Genres          []Genre  `json:"genres"`
	Authors         []Author `json:"authors"`
}

// Volume represents the Volumes table.
type Volume struct {
	VolumeID     int       `json:"volumeId"`
	LightNovelID int       `json:"lightNovelId"`
	VolumeNumber int       `json:"volumeNumber"`
	CoverArtURL  string    `json:"coverArtUrl"`
	ReleaseDate  time.Time `json:"releaseDate"`
}

// Genre represents the Genres table.
type Genre struct {
	GenreID int    `json:"genreId"`
	Name    string `json:"name"`
}

// LightNovelGenre represents the LightNovelGenres junction table for a Many-to-Many relationship between LightNovels and Genres.
type LightNovelGenre struct {
	LightNovelID int `json:"lightNovelId"`
	GenreID      int `json:"genreId"`
}

// Author represents the Authors table.
type Author struct {
	AuthorID int    `json:"authorId"`
	Name     string `json:"name"`
}

// LightNovelAuthor represents the LightNovelAuthors junction table for a Many-to-Many relationship between LightNovels and Authors.
type LightNovelAuthor struct {
	LightNovelID int `json:"lightNovelId"`
	AuthorID     int `json:"authorId"`
}

// User represents the Users table.
type User struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserList represents the UserLists table for tracking novels, their status, and progress.
type UserList struct {
	UserID   int    `json:"userId"`
	VolumeID int    `json:"volumeId"`
	Status   string `json:"status"`
	Progress int    `json:"progress"`
}
