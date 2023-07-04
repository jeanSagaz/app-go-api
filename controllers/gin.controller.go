package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.app.api/generics"
	"go.app.api/models"
)

func GinHandleRequests() {
	fmt.Println("Rest API v2.0 - gin Routers")

	router := gin.Default()
	router.GET("/albums", getAlbumsGin)
	router.GET("/albums/:id", getAlbumByIdGin)
	router.POST("/albums", postAlbumsGin)
	router.PUT("/albums/:id", putAlbumGin)
	router.DELETE("/albums/:id", deleteAlbumGin)

	log.Fatal(router.Run(":8080"), router)
}

func getAlbumsGin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByIdGin(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbumsGin(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbumGin(c *gin.Context) {
	var updateAlbum models.Album
	id := c.Param("id")

	// al := generics.FirstOrDefault(albums, func(p *album) bool { return p.Id == id })
	// if al == nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	// 	return
	// }

	idx := generics.Find(albums, func(value interface{}) bool {
		return value.(models.Album).Id == id
	})
	if idx < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}

	albums[idx] = updateAlbum
	c.IndentedJSON(http.StatusCreated, updateAlbum)
}

func deleteAlbumGin(c *gin.Context) {
	id := c.Param("id")

	al := generics.FirstOrDefault(albums, func(p *models.Album) bool { return p.Id == id })
	if al == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	fmt.Println(al)

	idx := generics.Find(albums, func(value interface{}) bool {
		return value.(models.Album).Id == id
	})
	if idx < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	//albums = RemoveIndex(albums, idx)
	albums = generics.FindAndDelete(albums, func(p *models.Album) bool { return p.Id == id })

	fmt.Println(albums)

	c.Status(http.StatusNoContent)
}
