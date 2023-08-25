package main

import (
	"log"

	config "github.com/AndrewAlizaga/CoolMusicApp/internal/config"
)

// The Main entry function connects into the configuration package and initialices the server
func main() {

	log.Println("RUNNING")
	config.InitServer()

}
