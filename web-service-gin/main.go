package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 10.99},
	{ID: "2", Title: "The Wall", Artist: "Pink Floyd", Price: 10.99},
	{ID: "3", Title: "The Division Bell", Artist: "Pink Floyd", Price: 10.99},
	{ID: "4", Title: "The Final Cut", Artist: "Pink Floyd", Price: 10.99},
	{ID: "5", Title: "Meddle", Artist: "Pink Floyd", Price: 10.99},
	{ID: "6", Title: "Wish You Were Here", Artist: "Pink Floyd", Price: 10.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	var id string
	id = c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.Run(":8080")
}
