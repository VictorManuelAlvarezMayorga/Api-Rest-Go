package controllers

import (
	"net/http"
	"strconv"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Cars)
}

func PostCars(c *gin.Context) {
	var newCar models.Car

	if err := c.BindJSON(&newCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Error al registrar Auto", "Detalles": err.Error()})
		return
	}

	models.Cars = append(models.Cars, newCar)

	c.IndentedJSON(http.StatusCreated, newCar)
}

func GetCarByID(c *gin.Context) {

	for _, a := range models.Cars {
		id, err := strconv.Atoi(c.Param("id")) // strcon convierte string a int
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
			return
		}

		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auto no encontrado"})
}

func DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range models.Cars {
		if a.ID == id {
			models.Cars = append(models.Cars[:i], models.Cars[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Eliminado con exito"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auto no encontrado"})
}

func EditCar(c *gin.Context) {

	var updatedCar models.Car
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	if err := c.BindJSON(&updatedCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al actualizar Auto",
			"detalle": err.Error(),
		})
		return
	}

	for i, a := range models.Cars {
		if a.ID == id {
			updatedCar.ID = a.ID
			models.Cars[i] = updatedCar
			c.IndentedJSON(http.StatusOK, models.Cars[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auto no encontrado"})
}
