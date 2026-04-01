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

	log.Printf("=== PORT value: '%s' ===", port) // ✅ Veremos el valor en logs

	if port == "" {
		port = "10000"
	}

	log.Printf("=== Iniciando en 0.0.0.0:%s ===", port)

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error:", err)
	}
}
