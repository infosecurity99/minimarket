package models

type Sale struct {
	ID               string `json:"id"`
	Branch_id        string `json:"branch_id"`
	Shopassistant_id string `json:"shopassistant_id"`
	Cashier_id       string `json:"cashier_id"`
	Payment_type     string `json:"payment_type"`
	Price            int    `json:"price"`
	Status_type      string `json:"status_type"`
	Clientname       string `json:"clientname"`
}
type CreateSale struct {
	Branch_id        string `json:"branch_id"`
	Shopassistant_id string `json:"shopassistant_id"`
	Cashier_id       string `json:"cashier_id"`
	Payment_type     string `json:"payment_type"`
	Price            int    `json:"price"`
	Status_type      string `json:"status_type"`
	Clientname       string `json:"clientname"`
}
type UpdateSale struct {
	ID               string `json:"id"`
	Branch_id        string `json:"branch_id"`
	Shopassistant_id string `json:"shopassistant_id"`
	Cashier_id       string `json:"cashier_id"`
	Price            int    `json:"price"`
	Clientname       string `json:"clientname"`
}

type SaleRepo struct {
	Sales []Sale `json:"sales"`
	Count int    `json:"count"`
}
