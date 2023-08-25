package config

import (
	"log"
	"os"

	routes "github.com/AndrewAlizaga/CoolMusicApp/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var baseUrl string = "/api/v1"

func InitServer() {
	godotenv.Load(".env")
	port := os.Getenv("APP_PORT")
	createRoutes(port)
}

func createRoutes(port string) {
	gin.SetMode(gin.ReleaseMode)
	log.Printf("Gin cold Start, PORT: %v \n", port)
	router := gin.New()
	router.Use(CORSMiddleware())
	routes.SetMusicRoutes(router, baseUrl)

	//TEST
	router.GET(baseUrl+"/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
		})
	})

	//Run Application
	router.Run(":" + port)
}

// CORSMiddleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, x-api-key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Next()
	}
}
