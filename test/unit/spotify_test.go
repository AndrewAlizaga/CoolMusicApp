package unit

import (
	"log"
	"testing"

	"github.com/AndrewAlizaga/CoolMusicApp/utils"
	"github.com/joho/godotenv"
)

func TestSpotifyConn(t *testing.T) {
	log.Println("init TestSpotifyGetTrackIRSC")
	godotenv.Load(".env")
	client := utils.GetSpotifyClient()
	log.Println(client)
	t.Log("success")

}
