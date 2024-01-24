package models

type Branch struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
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
