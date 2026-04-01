package main

import (
	"os"
	"ui2/database"
	"ui2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() //conexion a la db en la nube al iniciar

	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT") //render asigna el puerto automáticamente
	if port == "" {
		port = "8080" //si no hay variable, usa 8080 localmente por defecto
	}

	addr := "0.0.0.0:" + port
	router.Run(addr)

}
