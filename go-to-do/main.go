package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("what-a-secret")

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "todo-app",
		"aud": getRole(username),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	fmt.Printf("TokenClaims added: %+v\n", claims)
	return tokenString, nil
}

func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

type Todo struct {
	Text string
	Done bool
}

var todos []Todo
var loggedInUser string

func toggleIndex(index string) {
	i, _ := strconv.Atoi(index)
	if i >= 0 && i < len(todos) {
		todos[i].Done = !todos[i].Done
	}
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Todos":    todos,
			"LoggedIn": loggedInUser != "",
			"Username": loggedInUser,
		})
	})

	router.POST("/add", func(c *gin.Context) {
		text := c.PostForm("todo")
		todo := Todo{Text: text, Done: false}
		todos = append(todos, todo)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/toggle", func(c *gin.Context) {
		index := c.PostForm("index")
		toggleIndex(index)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
