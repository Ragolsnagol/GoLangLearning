package main

import (
	"example.com/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
