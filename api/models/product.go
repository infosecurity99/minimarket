package models

import (
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Barcode     string    `json:"barcode"`
	Category_id string    `json:"category_id"`
	Create_at   time.Time `json:"create_at"`
}

type CreateProduct struct {
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Category_id string `json:"category_id"`
	Barcode     int    `json:"barcode"`
}

type UpdateProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Category_id string `json:"category_id"`
}

type ProductResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}

type ProductSell struct {
	SelectedProducts  SellRequest    `json:"selected_products"`
	ProductPrices     map[string]int `json:"product_prices"`
	NotEnoughProducts map[string]int `json:"not_enough_products"`
	Prices            map[string]int `json:"prices"`
	NewProducts       map[string]int `json:"new_products"`
	ProductsBranchID  string         `json:"products_branch_id"`
}
type SellRequest struct {
	Products map[string]int `json:"products"`
	BasketID string         `json:"basket_id"`
	BranchID string         `json:"branch_id"`
}
