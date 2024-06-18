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

func authenticateMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	c.Next()
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
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
			"Role":     getRole(loggedInUser),
		})
	})

	router.POST("/add", authenticateMiddleware, func(c *gin.Context) {
		text := c.PostForm("todo")
		todo := Todo{Text: text, Done: false}
		todos = append(todos, todo)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/toggle", authenticateMiddleware, func(c *gin.Context) {
		index := c.PostForm("index")
		toggleIndex(index)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Dummy credential check
		if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
			tokenString, err := createToken(username)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error creating token")
				return
			}

			loggedInUser = username
			fmt.Printf("Token created: %s\n", tokenString)
			c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.String(http.StatusUnauthorized, "Invalid credentials")
		}
	})

	router.GET("/logout", func(c *gin.Context) {
		loggedInUser = ""
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
