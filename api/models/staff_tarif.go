package models

import "time"

type Staff_Tarif struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Tarif_Type_Enum  string    `json:"tarif_type"`
	Amount_For_Cashe int       `json:"amount_for_cashe"`
	Amount_For_Card  int       `json:"amount_for_card"`
	Create_at        time.Time `json:"ctreate_at"`
}

type CreateStaff_Tarif struct {
    Name             string `json:"name"`
    Tarif_Type_Enum  string `json:"tarif_type"`
    Amount_For_Cashe int    `json:"amount_for_cashe"`
    Amount_For_Card  int    `json:"amount_for_card"`
}


type UpdateStaff_Tarif struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Amount_For_Cashe int       `json:"amount_for_cashe"`
	Amount_For_Card  int       `json:"amount_for_card"`
	Create_at        time.Time `json:"ctreate_at"`
}

type Staff_Tarif_Repo struct {
	Staff_Tarif_Repos []Staff_Tarif `json:"staff_tarif_repos"`
	Count             int           `json:"count"`
}
