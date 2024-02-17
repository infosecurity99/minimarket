package models

import "time"

type Staff struct {
	ID         string    `json:"id"`
	Branch_id  string    `json:"branch_id"`
	Tarif_id   string    `json:"tarif_id"`
	Type_stuff string    `json:"type_stuff"`
	Name       string    `json:"name"`
	Balance    uint64    `json:"balance"`
	Age        int       `json:"age"`
	BirthDate  time.Time `json:"birthdate"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Create_at time.Time `json:"ctreate_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CreateStaff struct {
	Branch_id  string `json:"branch_id"`
	Tarif_id   string `json:"tarif_id"`
	Type_stuff string `json:"type_stuff"`
	Name       string `json:"name"`
	Balance    uint64 `json:"balance"`
	BirthDate  string `json:"birthdate"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type UpdateStaff struct {
	ID        string `json:"-"`
	BranchID  string `json:"branch_id"`
	TariffID  string `json:"tariff_id"`
	StaffType string `json:"staff_type"`
	Name      string `json:"name"`
	Balance   uint64 `json:"balance"`
	Login     string `json:"login"`
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
