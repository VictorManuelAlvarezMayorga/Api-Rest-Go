package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID      string `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Version string `json:"version"`
	Year    int    `json:"year"`
}

var cars = []car{
	{ID: "1", Brand: "Chevrolet", Model: "Camaro", Version: "RS", Year: 1967},
	{ID: "2", Brand: "Ford", Model: "Mustang", Version: "Boss 429", Year: 1969},
	{ID: "3", Brand: "Mazda", Model: "RX7", Version: "Coupé", Year: 1977},
	{ID: "4", Brand: "Dodge", Model: "Charger", Version: "Daytona", Year: 1970},
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func postCars(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Error al registrar Auto", "Detalles": err.Error()})
		return
	}

	cars = append(cars, newCar)

	c.IndentedJSON(http.StatusCreated, cars)

}

func getCarByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Auto no encontrado"})
}

func deleteCar(c *gin.Context) {
	id := c.Param("id")

	for i, a := range cars {
		if a.ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "Eliminado con exito"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Auto no encontrado"})
}

func editCar(c *gin.Context) {
	id := c.Param("id")
	var updatedCar car

	if err := c.BindJSON(&updatedCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al actualizar Auto",
			"detalle": err.Error(),
		})
		return
	}

	for i, a := range cars {
		if a.ID == id {
			updatedCar.ID = a.ID // Conserva el ID original
			cars[i] = updatedCar
			c.IndentedJSON(http.StatusOK, cars[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auto no encontrado"})
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", postCars)
	router.GET("/cars/:id", getCarByID)
	router.DELETE("/cars/:id", deleteCar)
	router.PUT("/cars/:id", editCar)

	router.Run("localhost:8080")
}
