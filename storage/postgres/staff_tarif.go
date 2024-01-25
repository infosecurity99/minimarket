package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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
func (s *staftarifRepo) CreateStaff_Tarifs(createStaff_tarif models.CreateStaff_Tarif) (string, error) {

	uid := uuid.New()
	create_at := time.Now()
	if _, err := s.db.Exec(`insert into 
	staff_tarif values ($1, $2, $3, $4, $5, $6)
			`,
		uid,
		createStaff_tarif.Name,
		createStaff_tarif.Tarif_Type_Enum,
		createStaff_tarif.Amount_For_Cashe,
		createStaff_tarif.Amount_For_Card,
		create_at,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyid  staftarif
func (s *staftarifRepo) GetByIdStaff_Tarifs(models.PrimaryKey) (models.Staff_Tarif, error) {
	return models.Staff_Tarif{}, nil
}

//get list
func (s *staftarifRepo) GetListStaff_Tarifs(models.GetListRequest) (models.Staff_Tarif_Repo, error) {
	return models.Staff_Tarif_Repo{}, nil
}

//update list
func (s *staftarifRepo) UpdateStaff_Tarifs(models.UpdateStaff_Tarif) (string, error) {
	return "", nil
}

//delete staf tarif
func (s *staftarifRepo) DeleteStaff_Tarifs(models.PrimaryKey) error {
	return nil
}
