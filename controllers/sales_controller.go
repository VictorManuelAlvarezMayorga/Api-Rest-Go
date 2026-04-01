package controllers

import (
	"net/http"
	"strconv"
	"ui2/models"

	"github.com/gin-gonic/gin"
)

var Sales = []models.Sale{
	{ID: 1, CarID: 1, CustomerID: 2, SaleDate: "2024-03-15", Price: 85000.00},
	{ID: 2, CarID: 3, CustomerID: 1, SaleDate: "2024-06-22", Price: 120000.00},
	{ID: 3, CarID: 4, CustomerID: 3, SaleDate: "2025-01-10", Price: 95000.00},
}

func GetSales(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Sales)
}

func GetSaleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for _, sale := range Sales {
		if sale.ID == id {
			c.IndentedJSON(http.StatusOK, sale)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Venta no encontrada"})
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

	Sales = append(Sales, newSale)
	c.IndentedJSON(http.StatusCreated, newSale)
}

func DeleteSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, sale := range Sales {
		if sale.ID == id {
			Sales = append(Sales[:i], Sales[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Venta eliminada correctamente"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Venta no encontrada"})
}
