package handlers

import (
	"net/http"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, db.Books)
}

func PostBook(c *gin.Context) {
	var newBook models.Book

	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}

	db.Books = append(db.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	for _, a := range db.Books {
		if a.ISBN == isbn {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func DeleteBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	for idx, a := range db.Books {
		if a.ISBN == isbn {
			c.JSON(http.StatusOK, a)
			db.Books = append(db.Books[:idx], db.Books[idx+1:]...)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func UpdateBookByISBN(c *gin.Context) {
	var updateBook models.Book
	isbn := c.Param("isbn")

	err := c.BindJSON(&updateBook)
	updateBook.ISBN = isbn
	if err != nil {
		return
	}

	for idx, a := range db.Books {
		if a.ISBN == isbn {
			db.Books[idx] = updateBook
			c.JSON(http.StatusOK, updateBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
