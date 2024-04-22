package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"example.com/ginapi/config"
	"example.com/ginapi/tutorial"
)

// getAuthors responds with the list of all authors as JSON.
func GetAuthors(c *gin.Context) {
	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, config.ConfigSettings.DbConnection)
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	authors, _ := queries.ListAuthors(ctx)

	c.IndentedJSON(http.StatusOK, authors)
}

// getAuthorByID locates the author whose ID value matches the id
// parameter sent by the client, then returns that author as a response.
func GetAuthorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, config.ConfigSettings.DbConnection)
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	fetchedAuthor, _ := queries.GetAuthor(ctx, int64(id))

	c.IndentedJSON(http.StatusOK, fetchedAuthor)
}

// postAuthors adds an author from JSON received in the request body.
func PostAuthor(c *gin.Context) {
	var newAuthor tutorial.Author

	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, config.ConfigSettings.DbConnection)
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// Call BindJSON to bind the received JSON to
	// newAuthor.
	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	insertedAuthor, _ := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: newAuthor.Name,
		Bio:  newAuthor.Bio,
	})

	c.IndentedJSON(http.StatusCreated, insertedAuthor)
}
