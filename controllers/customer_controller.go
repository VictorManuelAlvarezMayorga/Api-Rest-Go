package controllers

import (
	"net/http"
	"strconv"
	"ui2/database"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, first_name, last_name, email, phone, address FROM customers")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener clientes"})
		return
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.Address)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
			return
		}
		customers = append(customers, customer)
	}
	c.IndentedJSON(http.StatusOK, customers)
}

func GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var customer models.Customer
	err = database.DB.QueryRow("SELECT id, first_name, last_name, email, phone, address FROM customers WHERE id = $1", id).
		Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.Address)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente no encontrado"})
		return
	}
	c.IndentedJSON(http.StatusOK, customer)
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

	err := database.DB.QueryRow(
		"INSERT INTO customers (first_name, last_name, email, phone, address) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		newCustomer.FirstName, newCustomer.LastName, newCustomer.Email, newCustomer.Phone, newCustomer.Address,
	).Scan(&newCustomer.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar cliente"})
		return
	}
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

	_, err = database.DB.Exec(
		"UPDATE customers SET first_name=$1, last_name=$2, email=$3, phone=$4, address=$5 WHERE id=$6",
		updatedCustomer.FirstName, updatedCustomer.LastName, updatedCustomer.Email, updatedCustomer.Phone, updatedCustomer.Address, id,
	)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar cliente"})
		return
	}

	updatedCustomer.ID = id
	c.IndentedJSON(http.StatusOK, updatedCustomer)
}

func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar cliente"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Cliente eliminado correctamente"})
}
