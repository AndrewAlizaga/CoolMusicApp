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

// The createRoutes internal funciton configures Gin
// it also subgroups the routes for music management
func createRoutes(port string) {
	gin.SetMode(gin.ReleaseMode)
	log.Printf("Gin cold Start, PORT: %v \n", port)
	router := gin.New()
	router.Use(CORSMiddleware())
	routes.SetMusicRoutes(router, baseUrl)

	router.GET(baseUrl+"/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
		})
	})

	//Run Gin
	router.Run(":" + port)
}

// CORSMiddleware
// the CORSMiddleware function configures management for Access controll, application allowed methods, and headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, x-api-key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Next()
	}
}
