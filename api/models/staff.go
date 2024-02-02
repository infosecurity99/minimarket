package models

import "time"

type Staff struct {
	ID         string    `json:"id"`
	Branch_id  string    `json:"branch_id"`
	Tarif_id   string    `json:"tarif_id"`
	Type_stuff string    `json:"type_stuff"`
	Name       string    `json:"name"`
	Balance    string    `json:"balance"`
	Age        int       `json:"age"`
	BirthDate  string    `json:"birthdate"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Create_at  time.Time `json:"ctreate_at"`
}

type CreateStaff struct {
	Branch_id  string `json:"branch_id"`
	Tarif_id   string `json:"tarif_id"`
	Type_stuff string `json:"type_stuff"`
	Name       string `json:"name"`
	Balance    string `json:"balance"`
	BirthDate  string `json:"birthdate"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type UpdateStaff struct {
	ID        string `json:"id"`
	Branch_id string `json:"branch_id"`
	Tarif_id  string `json:"tarif_id"`
	Name      string `json:"name"`
	Balance   string `json:"balance"`
	Age       int    `json:"age"`
	BirthDate string `json:"birthdate"`
}

type StaffRepo struct {
	Staffs []Staff `json:"staffs"`
	Count  int     `json:"count"`
}

type UpdateStaffPassword struct {
	ID          string `json:"-"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
