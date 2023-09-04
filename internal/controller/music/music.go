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

// Controller for posting track from Spotify by ISRC
func PostMusic(c *gin.Context) {

	var body *isrc.ISRC = &isrc.ISRC{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "NOT OK, MISSING ISRC"})
		return
	}

	spotifyClient := utils.GetSpotifyClient()
	spotifyTrack, err := spotifyClient.Search("isrc:"+body.Isrc, spotify.SearchTypeTrack)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"result": "TRACK NOT FOUND - 💀"})
		return
	}

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
	if err := dbConn.Create(&Track).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"result": "error - possible duplicated"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": Track})
	return
}

// Controller for Find By Match route
func FindByMatch(c *gin.Context) {

	match := c.Query("match")

	if match == "" {
		c.JSON(http.StatusBadRequest, gin.H{"result": "MISSING ISRC"})
		return
	}

	connDB := orm.GetDB()
	var tracks []sqlModels.Track

	searchPattern := "%" + match + "%"

	txResult := connDB.Where("Title LIKE ?", searchPattern).Find(&tracks)

	if txResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"results": txResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": tracks})
	return
}

// Controller for Get By ISRC route
func GetByISRC(c *gin.Context) {

	isrc := c.Param("id")

	if isrc == "" || isrc == ":id" {
		c.JSON(http.StatusBadRequest, gin.H{"result": "MISSING ISRC"})
		return
	}

	connDB := orm.GetDB()

	var track sqlModels.Track
	txResult := connDB.First(&track, "isrc = ?", isrc)

	if txResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"results": txResult.Error.Error()})
		return
	}

	//TODO: SQL LOGIC
	c.JSON(http.StatusOK, gin.H{"result": track})
	return
}
