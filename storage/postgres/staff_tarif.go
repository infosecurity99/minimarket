package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
)

type staftarifRepo struct {
	db *sql.DB
}

func NewStaff_Tarif(db *sql.DB) storage.IStaff_Tarif {
	return &staftarifRepo{
		db: db,
	}
}

//create stafftarif
func (s staftarifRepo) CreateStaff_Tarif(models.CreateStaff_Tarif) (string, error) {
	return "", nil
}

//getbyid  staftarif

func (s staftarifRepo) GetByIdStaff_Tarif(models.PrimaryKey) (models.Staff_Tarif, error) {
	return models.Staff_Tarif{}, nil
}

//get list

func (s staftarifRepo) GetListStaff_Tarif(models.GetListRequest) (models.Staff_Tarif_Repo, error) {
	return models.Staff_Tarif_Repo{}, nil
}

//update list

func (s staffRepo) UpdateStaff_Tarif(models.Staff_Tarif) (string, error) {
	return "", nil
}

//delete staf tarif

func (s staffRepo) DeleteStaff_Tarif(models.PrimaryKey) error {
	return nil
}
