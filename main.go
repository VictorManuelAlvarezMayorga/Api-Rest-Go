package main

import (
	"os"
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT") //render asigna el puerto automáticamente
	if port == "" {
		port = "8080" //si no hay variable, usa 8080 localmente por defecto
	}

	router.Run(":" + port)
}
