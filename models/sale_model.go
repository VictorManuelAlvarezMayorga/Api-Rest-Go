package models

type Sale struct {
	ID         int     `json:"id"`
	CarID      int     `json:"car_id"`
	CustomerID int     `json:"customer_id"`
	SaleDate   string  `json:"sale_date"`
	Price      float64 `json:"price"`
}

// SaleDetail es lo que devuelve el GET con los datos relacionados
type SaleDetail struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Brand     string  `json:"brand"`
	Model     string  `json:"model"`
	Version   string  `json:"version"`
	SaleDate  string  `json:"sale_date"`
	Price     float64 `json:"price"`
}
