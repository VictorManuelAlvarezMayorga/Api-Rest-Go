package controllers

import (
	"net/http"
	"strconv"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

var Customers = []models.Customer{
	{ID: 1, FirstName: "Fercho", LastName: "Urquiza", Email: "fergurquiza@gmail.com", Phone: "4491234567", Address: "Av. López Mateos 123, Aguascalientes"},
	{ID: 2, FirstName: "Ana", LastName: "González", Email: "anagabycars@gmail.com", Phone: "4499876543", Address: "Calle Morelos 456, Guadalajara"},
	{ID: 3, FirstName: "Roberto", LastName: "Musso", Email: "roberto.musso@gmail.com", Phone: "4495556789", Address: "Blvd. Zacatecas 789, Aguascalientes"},
}

func GetCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Customers)
}

func GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for _, customer := range Customers {
		if customer.ID == id {
			c.IndentedJSON(http.StatusOK, customer)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente no encontrado"})
}

func PostCustomer(c *gin.Context) {
	var newCustomer models.Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al registrar cliente",
			"detalle": err.Error(),
		})
		return
	}

	Customers = append(Customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

func UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedCustomer models.Customer

	if err := c.BindJSON(&updatedCustomer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al actualizar cliente",
			"detalle": err.Error(),
		})
		return
	}

	for i, customer := range Customers {
		if customer.ID == id {
			updatedCustomer.ID = customer.ID
			Customers[i] = updatedCustomer
			c.IndentedJSON(http.StatusOK, Customers[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente no encontrado"})
}

func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, customer := range Customers {
		if customer.ID == id {
			Customers = append(Customers[:i], Customers[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Cliente eliminado correctamente"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente no encontrado"})
}
