package main

import (
	"fmt"
	"net/http"
	"os"
	"ui2/database"
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Iniciando servidor en puerto:", port)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: router,
	}

	srv.ListenAndServe()
}
