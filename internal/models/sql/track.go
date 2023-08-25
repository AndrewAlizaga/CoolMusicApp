package sql

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Title      string
	ArtistName string
	Popularity int32
	Image      string
	Isrc       string `gorm:"unique"`
}
