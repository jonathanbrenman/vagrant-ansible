package main

import (
	"app1/api/router"
)

func main() {
	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}
