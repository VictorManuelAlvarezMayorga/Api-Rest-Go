package main

import (
	"log"
	"net/http"
	"os"
	"ui2/database"
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: router,
	}

	log.Printf("Servidor iniciando en 0.0.0.0:%s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
