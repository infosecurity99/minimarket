package models

import "time"

type Branch struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Create_at time.Time `json:"ctreate_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CreateBranch struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type UpdateBranch struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchResponse struct {
	Branches []Branch `json:"branches"`
	Count    int      `json:"count"`
}
