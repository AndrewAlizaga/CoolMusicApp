package orm

import (
	"log"

	"github.com/AndrewAlizaga/CoolMusicApp/internal/models/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {

	log.Println("invoke GetDB")

	if DB == nil {
		log.Println("initing db")
		var err error
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate Schema
		DB.AutoMigrate(&sql.Track{})
		DB.Exec("CREATE UNIQUE INDEX idx_unique_track_irsc ON tracks (id)")

	}

	return DB
}
