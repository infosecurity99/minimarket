package models

import "time"

type TransactionStorage struct {
	ID               string    `json:"id"`
	Branch_id        string    `json:"branch_id"`
	Staff_id         string    `json:"staff_id"`
	Product_id       string    `json:"product_id"`
	Transaction_type string    `json:"transaction_type"`
	Price            uint64   `json:"price"`
	Quantity         uint64   `json:"quantity"`
	Create_at time.Time `json:"ctreate_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CreateTransactionStorage struct {
	Branch_id        string  `json:"branch_id"`
	Staff_id         string  `json:"staff_id"`
	Product_id       string  `json:"product_id"`
	Transaction_type string  `json:"transaction_type"`
	Price            uint64 `json:"price"`
	Quantity         uint64 `json:"quantity"`
}

type UpdateTransactionStorage struct {
	ID         string  `json:"id"`
	Branch_id  string  `json:"branch_id"`
	Staff_id   string  `json:"staff_id"`
	Product_id string  `json:"product_id"`
	Price      uint64 `json:"price"`
	Quantity   uint64 `json:"quantity"`
}

type TransactionStorageResponse struct {
	TransactionStorages []TransactionStorage `json:"transactionstorages"`
	Count               int                  `json:"count"`
}
