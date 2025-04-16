package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Doe"},
	{ID: "2", Title: "Jeru", Artist: "Gerry"},
	{ID: "3", Title: "Sarah Vaughan and Clifford"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	for _, album := range albums {
		if album.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, album)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func deleteAlbums(c *gin.Context) {
	for _, album := range albums {
		if album.ID == c.Param("id") {
			c.IndentedJSON(http.StatusNoContent, gin.H{})
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			newAlbum.ID = id
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Álbum no encontrado"})
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.DELETE("/albums/:id", deleteAlbums)
	router.PUT("/albums/:id", updateAlbum)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
