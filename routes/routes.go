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
	// Customers
	router.GET("/customers", controllers.GetCustomers)
	router.POST("/customers", controllers.PostCustomer)
	router.GET("/customers/:id", controllers.GetCustomerByID)
	router.PUT("/customers/:id", controllers.UpdateCustomer)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)

	// Sales
	router.GET("/sales", controllers.GetSales)
	router.POST("/sales", controllers.PostSale)
	router.GET("/sales/:id", controllers.GetSaleByID)
	router.DELETE("/sales/:id", controllers.DeleteSale)

	router.Run("localhost:8080")
}
