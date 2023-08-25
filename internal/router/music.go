package router

import (
	musicController "github.com/AndrewAlizaga/CoolMusicApp/internal/controller/music"
	"github.com/gin-gonic/gin"
)

// Music routes group
func SetMusicRoutes(router *gin.Engine, baseUrl string) {

	musicGroups := router.Group(baseUrl + "/music")
	musicGroups.POST("/", musicController.PostMusic)
	musicGroups.GET("/search", musicController.FindByMatch)
	musicGroups.GET("/isrc/:id", musicController.GetByISRC)

}
