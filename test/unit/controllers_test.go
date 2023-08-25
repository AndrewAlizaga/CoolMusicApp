package unit

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	musicController "github.com/AndrewAlizaga/CoolMusicApp/internal/controller/music"
	modelIsrc "github.com/AndrewAlizaga/CoolMusicApp/internal/models/isrc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func init() {
	//remove db for operations
	os.Remove("test.db")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostMusicControllerSuccess(t *testing.T) {

	r := SetUpRouter()
	r.POST("/api/v1/music", musicController.PostMusic)

	body := &modelIsrc.ISRC{
		Isrc: "GBAYE0601498",
	}

	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/api/v1/music", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Println("err: ", err.Error())
		t.Logf(err.Error())
		t.Fail()
	}

	log.Println("err: ", err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostMusicControllerFail(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/v1/music", musicController.PostMusic)

	body := &modelIsrc.ISRC{
		Isrc: "GBAYE0601498",
	}

	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/api/v1/music", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Println("err: ", err.Error())
		t.Logf(err.Error())
		t.Fail()
	}

	log.Println("err: ", err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestPostMusicControllerGetByISRC(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/music/isrc/:id", musicController.GetByISRC)

	req, err := http.NewRequest("GET", "/api/v1/music/isrc/GBAYE0601498", nil)

	if err != nil {
		log.Println("err: ", err.Error())
		t.Logf(err.Error())
		t.Fail()
	}

	log.Println("err: ", err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostMusicControllerGetMany(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/music/search", musicController.FindByMatch)

	req, err := http.NewRequest("GET", "/api/v1/music/search?match=yell", nil)

	if err != nil {
		log.Println("err: ", err.Error())
		t.Logf(err.Error())
		t.Fail()
	}

	log.Println("err: ", err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
