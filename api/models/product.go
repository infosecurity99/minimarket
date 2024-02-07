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
	Price       float64   `json:"price"`
	Category_id string `json:"category_id"`
	Barcode     int    `json:"barcode"`
}

type UpdateProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       float64   `json:"price"`
	Category_id string `json:"category_id"`
}

type ProductResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}

