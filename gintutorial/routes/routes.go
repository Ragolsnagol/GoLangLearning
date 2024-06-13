package routes

import (
	"fmt"
	"time"

	"example.com/handlers"
	"github.com/gin-gonic/gin"
)

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(requestLogger())

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:isbn", handlers.GetBookByISBN)
	router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
	router.PUT("/books/:isbn", handlers.UpdateBookByISBN)
	router.POST("/books", handlers.PostBook)

	return router
}
