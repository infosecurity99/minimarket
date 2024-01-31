package models

import (
	"time"
)

type Basket struct {
	ID         string    `json:"id"`
	Sale_id    string    `json:"sale_id"`
	Product_id string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	Create_at  time.Time `json:"ctreate_at"`
}

type CreateBasket struct {
	Sale_id    string    `json:"sale_id"`
	Product_id string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	
}

type UpdateBasket struct {
	ID         string    `json:"id"`
	Sale_id    string    `json:"sale_id"`
	Product_id string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`

}

type BasketResponse struct {
	Baskets []Basket `json:"baskets"`
	Count   int      `json:"count"`
}
