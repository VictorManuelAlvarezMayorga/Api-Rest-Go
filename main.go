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

	// Conexión a la base de datos
	database.Connect()
	log.Println("=== DB CONECTADA ===")

	// Configuración de Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Ruta base (IMPORTANTE para Render)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API funcionando correctamente",
		})
	})

	// Rutas de tu app
	routes.Routes(router)

	// Puerto dinámico (Render lo asigna)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" // fallback local
	}

	log.Printf("=== Servidor corriendo en puerto %s ===", port)

	// 🚀 ESTE es el cambio clave
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
