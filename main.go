package main

import (
	"log"

	config "github.com/AndrewAlizaga/CoolMusicApp/internal/config"
)

func main() {

	log.Println("RUNNING")
	config.InitServer()

}
