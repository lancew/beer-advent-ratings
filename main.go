package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type Beer struct {
	Alcohol_percentage float32 `json:"alcohol_percentage"`
	Date               string  `json:"date"`
	Name               string  `json:"name"`
	Notes              string  `json:"notes"`
	Rating             int     `json:"rating"`
	Beer_type          string  `json:"beer_type"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.LoadHTMLGlob("templates/*.html")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/beers", func(c *gin.Context) {

		content, err := ioutil.ReadFile("./beers.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		var beers []Beer
		err = json.Unmarshal(content, &beers)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		c.JSON(http.StatusOK, beers)

	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
