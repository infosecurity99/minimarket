package models

import "time"

type Sale struct {
	ID               string    `json:"id"`
	Branch_id        string    `json:"branch_id"`
	Shopassistant_id string    `json:"shopassistant_id"`
	Cashier_id       string    `json:"cashier_id"`
	Payment_type     string    `json:"payment_type"`
	Price            float64   `json:"price"`
	Status_type      string    `json:"status_type"`
	Clientname       string    `json:"clientname"`
	Create_at        time.Time `json:"ctreate_at"`
}
type CreateSale struct {
	Branch_id        string  `json:"branch_id"`
	Shopassistant_id string  `json:"shopassistant_id"`
	Cashier_id       string  `json:"cashier_id"`
	Payment_type     string  `json:"payment_type"`
	Price            float64 `json:"price"`
	Status_type      string  `json:"status_type"`
	Clientname       string  `json:"clientname"`
}
type UpdateSale struct {
	ID               string  `json:"id"`
	Branch_id        string  `json:"branch_id"`
	Shopassistant_id string  `json:"shopassistant_id"`
	Cashier_id       string  `json:"cashier_id"`
	Price            float64 `json:"price"`
	Clientname       string  `json:"clientname"`
}

type SaleRepos struct {
	Sales []Sale `json:"sales"`
	Count int    `json:"count"`
}

type StartSale struct {
	Branch_id        string `json:"branch_id"`
	Shopassistant_id string `json:"shopassistant_id"`
	Cashier_id       string `json:"cashier_id"`
	Status_type      string `json:"status_type"`
	Clientname       string `json:"clientname"`
	Payment_type     string `json:"payment_type"`
}

type EndSales struct {

}
