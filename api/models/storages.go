package models

import "time"

type Storage struct {
	ID         string    `json:"id"`
	Product_id string    `json:"product_id"`
	Branch_id  string    `json:"branch_id"`
	Count      int       `json:"count"`
	Create_at  time.Time `json:"ctreate_at"`
}

type CreateStorage struct {
	Product_id string    `json:"product_id"`
	Branch_id  string    `json:"branch_id"`
	Count      int       `json:"count"`

}

type UpdateStorage struct {
	ID         string `json:"id"`
	Product_id string `json:"product_id"`
	Branch_id  string `json:"branch_id"`
	Count      int    `json:"count"`
}

type StorageRepos struct {
	Storages []Storage `json:"storages"`
	Count    int       `json:"count"`
}
