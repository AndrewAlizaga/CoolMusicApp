package music

import (
	"log"
	"net/http"

	orm "github.com/AndrewAlizaga/CoolMusicApp/internal/orm"
	"github.com/AndrewAlizaga/CoolMusicApp/utils"
	"github.com/zmb3/spotify"

	"github.com/AndrewAlizaga/CoolMusicApp/internal/models/isrc"
	sqlModels "github.com/AndrewAlizaga/CoolMusicApp/internal/models/sql"

	"github.com/gin-gonic/gin"
)

func PostMusic(c *gin.Context) {

	var body *isrc.ISRC = &isrc.ISRC{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "NOT OK, MISSING ISRC"})
	}

	//TODO: SPOTIFY LOGIC
	spotifyClient := utils.GetSpotifyClient()

	if spotifyTrack, err := spotifyClient.Search("isrc:"+body.Isrc, spotify.SearchTypeTrack); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"result": "TRACK NOT FOUND - 💀"})
	} else {

		//TODO: SQL LOGIC

		if len(spotifyTrack.Tracks.Tracks) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"result": "NOTHING FOUND"})
			return
		}

		mostPopular := map[string]interface{}{
			"idx": 0,
			"lvl": 0,
		}

		for idx, track := range spotifyTrack.Tracks.Tracks {
			if track.Popularity >= mostPopular["lvl"].(int) {
				log.Println("replacing by popularity")
				mostPopular["idx"] = idx
				mostPopular["lvl"] = track.Popularity
			}
		}

		mostPopularTrack := spotifyTrack.Tracks.Tracks[mostPopular["idx"].(int)]

		var Track sqlModels.Track = sqlModels.Track{
			Isrc:       body.Isrc,
			Title:      mostPopularTrack.Name,
			ArtistName: mostPopularTrack.Artists[0].Name,
			Popularity: int32(mostPopularTrack.Popularity),
			Image:      mostPopularTrack.Album.Images[0].URL,
		}

		dbConn := orm.GetDB()
		err := dbConn.Create(&Track).Error

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "error - possible duplicated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": Track})
	}

}

func FindAll(c *gin.Context) {

	match := c.Query("match")

	if match == "" {
		c.JSON(http.StatusBadRequest, gin.H{"result": "NOT OK, MISSING ISRC"})

	}

	connDB := orm.GetDB()
	var tracks []sqlModels.Track

	searchPattern := "%" + match + "%"

	log.Println(searchPattern)
	connDB.Where("Title LIKE ?", searchPattern).Find(&tracks)

	c.JSON(http.StatusOK, gin.H{"result": tracks})
}

func GetByISRC(c *gin.Context) {

	isrc := c.Param("id")

	if isrc == "" {
		c.JSON(http.StatusBadRequest, gin.H{"result": "NOT OK, MISSING ISRC"})
	}

	connDB := orm.GetDB()

	var track sqlModels.Track
	connDB.First(&track, "isrc = ?", isrc)

	//TODO: SQL LOGIC
	c.JSON(http.StatusOK, gin.H{"result": track})
}
