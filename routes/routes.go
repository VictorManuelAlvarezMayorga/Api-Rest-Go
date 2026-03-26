package routes

import (
	"ui2/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/cars", controllers.GetCars)
	router.POST("/cars", controllers.PostCars)
	router.GET("/cars/:id", controllers.GetCarByID)
	router.DELETE("/cars/:id", controllers.DeleteCar)
	router.PUT("/cars/:id", controllers.EditCar)

	router.Run("localhost:8080")
}
