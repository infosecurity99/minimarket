package models

import "time"

type Transaction struct {
	ID                    string    `json:"id"`
	Sale_id               string    `json:"sale_id"`
	Staff_id              string    `json:"staff_id"`
	Transaction_type_enum string    `json:"transaction_type_enum"`
	Source_type_enum      string    `json:"source_type_enum"`
	Amount                string    `json:"amount"`
	Description           string    `json:"description"`
	Create_at             time.Time `json:"ctreate_at"`
}
type CreateTransaction struct {
	Sale_id               string    `json:"sale_id"`
	Staff_id              string    `json:"staff_id"`
	Transaction_type_enum string    `json:"transaction_type_enum"`
	Source_type_enum      string    `json:"source_type_enum"`
	Amount                string    `json:"amount"`
	Description           string    `json:"description"`
	Create_at             time.Time `json:"ctreate_at"`
}

type UpdateTransaction struct {
	Sale_id     string    `json:"sale_id"`
	Staff_id    string    `json:"staff_id"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Create_at   time.Time `json:"ctreate_at"`
}

type TransactionRepo struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}
