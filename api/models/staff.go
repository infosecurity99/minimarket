package models

type Staff struct {
	ID              string `json:"id"`
	Branch_id       string `json:"branch_id"`
	Tarif_id        string `json:"tarif_id"`
	Type_Stuff_Enum string `json:"type_stuff_enum"`
	Name            string `json:"name"`
	Balance         int    `json:"balance"`
	Age             int    `json:"age"`
	BirthDate       int    `json:"birthdate"`
	Login           string `json:"login"`
	Password        string `json:"password"`
}

type CreateStaff struct {
	Branch_id       string `json:"branch_id"`
	Tarif_id        string `json:"tarif_id"`
	Type_Stuff_Enum string `json:"type_stuff_enum"`
	Name            string `json:"name"`
	Balance         int    `json:"balance"`
	Age             int    `json:"age"`
	BirthDate       int    `json:"birthdate"`
	Login           string `json:"login"`
	Password        string `json:"password"`
}

type UpdateStaff struct {
	Branch_id string `json:"branch_id"`
	Tarif_id  string `json:"tarif_id"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	Age       int    `json:"age"`
	BirthDate int    `json:"birthdate"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type StaffRepo struct {
	Staffs []Staff `json:"staffs"`
	Count  int     `json:"count"`
}
