package controllers

import (
	"net/http"
	"strconv"
	"ui2/database"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, brand, model, version, year, image_url FROM cars WHERE isdeleted = false")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener autos"})
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Version, &car.Year, &car.ImageURL)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
			return
		}
		cars = append(cars, car)
	}
	c.IndentedJSON(http.StatusOK, cars)
}

func GetCarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var car models.Car
	err = database.DB.QueryRow("SELECT id, brand, model, version, year, image_url FROM cars WHERE id = $1 AND isdeleted = false", id).
		Scan(&car.ID, &car.Brand, &car.Model, &car.Version, &car.Year, &car.ImageURL)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auto no encontrado"})
		return
	}
	c.IndentedJSON(http.StatusOK, car)
}

func PostCars(c *gin.Context) {
	var newCar models.Car

	if err := c.BindJSON(&newCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al registrar Auto",
			"detalle": err.Error(),
		})
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO cars (brand, model, version, year, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		newCar.Brand, newCar.Model, newCar.Version, newCar.Year, newCar.ImageURL,
	).Scan(&newCar.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar auto"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newCar)
}

func EditCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedCar models.Car
	if err := c.BindJSON(&updatedCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al actualizar Auto",
			"detalle": err.Error(),
		})
		return
	}

	_, err = database.DB.Exec(
		"UPDATE cars SET brand=$1, model=$2, version=$3, year=$4, image_url=$5 WHERE id=$6",
		updatedCar.Brand, updatedCar.Model, updatedCar.Version, updatedCar.Year, updatedCar.ImageURL, id,
	)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar auto"})
		return
	}

	updatedCar.ID = id
	c.IndentedJSON(http.StatusOK, updatedCar)
}

func DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	_, err = database.DB.Exec("UPDATE cars SET isdeleted = true WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar auto"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Auto eliminado correctamente"})
}
