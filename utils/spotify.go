package utils

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var client spotify.Client

func GetSpotifyClient() spotify.Client {
	if client == (spotify.Client{}) {

		godotenv.Load(".env")
		authConfig := &clientcredentials.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_KEY"),
			TokenURL:     spotify.TokenURL,
		}

		accessToken, err := authConfig.Token(context.Background())
		if err != nil {
			log.Fatalf("error retrieve access token: %v", err)
		}

		client = spotify.Authenticator{}.NewClient(accessToken)
	}

	return client

}
