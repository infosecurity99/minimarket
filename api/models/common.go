package models

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}
type GetListRequestTransaction struct {
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
	Search     string  `json:"search"`
	ToAmount   float64 `json:"toamount"`
	FromAmount float64 `json:"fromamount"`
}
type GetListRequestSale struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	Search    string  `json:"search"`
	ToPrice   float64 `json:"toprice"`
	FromPrice float64 `json:"fromprice"`
}
