package main

import (
	"example.com/ginapi/api"
	"example.com/ginapi/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ReadConfig()

	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbums)
	router.GET("/authors", api.GetAuthors)
	router.GET("/authors/:id", api.GetAuthorByID)
	router.POST("/authors", api.PostAuthor)

	router.Run("localhost:8080")
}
