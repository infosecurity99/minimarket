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
	ToAmount   uint64 `json:"toamount"`
	FromAmount uint64 `json:"fromamount"`
}
type GetListRequestSale struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	Search    string  `json:"search"`
	ToPrice   uint64 `json:"toprice"`
	FromPrice uint64 `json:"fromprice"`
}
