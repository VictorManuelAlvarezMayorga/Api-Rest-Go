package main

import (
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	router.Run("localhost:8080")
}
