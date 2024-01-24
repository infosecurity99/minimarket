package models

type Transaction struct {
	ID                    string `json:"id"`
	Sale_id               string `json:"sale_id"`
	Staff_id              string `json:"staff_id"`
	Transaction_type_enum string `json:"transaction_type_enum"`
	Source_type_enum      string `json:"source_type_enum"`
	Amount                string `json:"amount"`
	Description           string `json:"description"`
}
type CreateTransaction struct {
	Sale_id               string `json:"sale_id"`
	Staff_id              string `json:"staff_id"`
	Transaction_type_enum string `json:"transaction_type_enum"`
	Source_type_enum      string `json:"source_type_enum"`
	Amount                string `json:"amount"`
	Description           string `json:"description"`
}

type UpdateTransaction struct {
	Sale_id     string `json:"sale_id"`
	Staff_id    string `json:"staff_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

type TransactionRepo struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}
