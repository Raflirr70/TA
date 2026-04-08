package main

import (
	"tipes/internal/config"
	"tipes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connection()

	r := gin.Default()
	routes.MainRoute(r)

	r.Run(":8080")
}
