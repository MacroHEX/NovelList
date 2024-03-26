package models

import "time"

type LightNovel struct {
	LightNovelID    int      `gorm:"primaryKey;column:lightnovelid" json:"lightNovelId"`
	Title           string   `gorm:"column:title" json:"title"`
	NumberOfVolumes int      `gorm:"column:numberofvolumes" json:"numberOfVolumes"`
	CoverArtURL     string   `gorm:"column:coverarturl" json:"coverArtUrl"`
	Volumes         []Volume `gorm:"foreignKey:LightNovelID" json:"volumes"`
	Genres          []Genre  `gorm:"many2many:lightnovelgenres;joinForeignKey:lightnovelid;joinReferences:genreid" json:"genres"`
	Authors         []Author `gorm:"many2many:lightnovelauthors;joinForeignKey:lightnovelid;joinReferences:authorid" json:"authors"`
}

func (LightNovel) TableName() string {
	return "lightnovels"
}

type Volume struct {
	VolumeID     int       `gorm:"primaryKey;column:volumeid" json:"volumeId"`
	LightNovelID int       `gorm:"column:lightnovelid" json:"lightNovelId"`
	VolumeNumber int       `gorm:"column:volumenumber" json:"volumeNumber"`
	CoverArtURL  string    `gorm:"column:coverarturl" json:"coverArtUrl"`
	ReleaseDate  time.Time `gorm:"column:releasedate" json:"releaseDate"`
}

type Genre struct {
	GenreID int    `gorm:"primaryKey;column:genreid" json:"genreId"`
	Name    string `gorm:"column:name" json:"name"`
}

type Author struct {
	AuthorID int    `gorm:"primaryKey;column:authorid" json:"authorId"`
	Name     string `gorm:"column:name" json:"name"`
}

type LightNovelAuthor struct {
	LightNovelID int `json:"lightNovelId"`
	AuthorID     int `json:"authorId"`
}
