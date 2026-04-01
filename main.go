package main

import (
	"log"
	"os"
	"ui2/database"
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("=== INICIANDO APP ===")

	database.Connect()
	log.Println("=== CONTINÚA APP (aunque falle DB) ===")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Ruta base obligatoria
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API funcionando",
		})
	})

	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	log.Println("🚀 Servidor iniciando en puerto", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
