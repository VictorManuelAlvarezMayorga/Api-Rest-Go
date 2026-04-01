package controllers

import (
	"net/http"
	"strconv"
	"ui2/database"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT s.id, c.first_name, c.last_name, ca.brand, ca.model, ca.version, s.sale_date, s.price
		FROM sales s
		JOIN customers c ON s.customer_id = c.id
		JOIN cars ca ON s.car_id = ca.id
		WHERE s.isdeleted = false
	`)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ventas"})
		return
	}
	defer rows.Close()

	var sales []models.SaleDetail
	for rows.Next() {
		var sale models.SaleDetail
		err := rows.Scan(&sale.ID, &sale.FirstName, &sale.LastName, &sale.Brand, &sale.Model, &sale.Version, &sale.SaleDate, &sale.Price)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
			return
		}
		sales = append(sales, sale)
	}
	c.IndentedJSON(http.StatusOK, sales)
}

func GetSaleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var sale models.SaleDetail
	err = database.DB.QueryRow(`
		SELECT s.id, c.first_name, c.last_name, ca.brand, ca.model, ca.version, s.sale_date, s.price
		FROM sales s
		JOIN customers c ON s.customer_id = c.id
		JOIN cars ca ON s.car_id = ca.id
		WHERE s.id = $1 AND s.isdeleted = false
	`, id).Scan(&sale.ID, &sale.FirstName, &sale.LastName, &sale.Brand, &sale.Model, &sale.Version, &sale.SaleDate, &sale.Price)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Venta no encontrada"})
		return
	}
	c.IndentedJSON(http.StatusOK, sale)
}

func PostSale(c *gin.Context) {
	var newSale models.Sale

	if err := c.BindJSON(&newSale); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Error al registrar venta",
			"detalle": err.Error(),
		})
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO sales (car_id, customer_id, sale_date, price) VALUES ($1, $2, $3, $4) RETURNING id",
		newSale.CarID, newSale.CustomerID, newSale.SaleDate, newSale.Price,
	).Scan(&newSale.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar venta"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newSale)
}

func DeleteSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	_, err = database.DB.Exec("UPDATE sales SET isdeleted = true, updatedat = now() WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar venta"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Venta eliminada correctamente"})
}
