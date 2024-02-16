package models

import "time"

type Transaction struct {
	ID               string    `json:"id"`
	Sale_id          string    `json:"sale_id"`
	Staff_id         string    `json:"staff_id"`
	Transaction_type string    `json:"transaction_type"`
	Sourcetype       string    `json:"sourcetype"`
	Amount           float64    `json:"amount"`
	Description      string    `json:"description"`
	Create_at        time.Time `json:"ctreate_at"`
}
type CreateTransaction struct {
	Sale_id          string `json:"sale_id"`
	Staff_id         string `json:"staff_id"`
	Transaction_type string `json:"transaction_type"`
	Sourcetype       string `json:"sourcetype"`
	Amount           float64 `json:"amount"`
	Description      string `json:"description"`
}

type UpdateTransaction struct {
	ID          string `json:"id"`
	Sale_id     string `json:"sale_id"`
	Staff_id    string `json:"staff_id"`
	Amount      float64 `json:"amount"`
	Description string `json:"description"`
}

type TransactionRepo struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}
