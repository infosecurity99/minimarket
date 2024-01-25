package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
)

type staffRepo struct {
	db *sql.DB
}

func NewStaffRepo(db *sql.DB) storage.IStaff {
	return &staffRepo{
		db: db,
	}
}

//create  for staff
func (s *staffRepo) CreateStaff(models.CreateStaff) (string, error) {
	return "", nil
}

//get by id for staff

func (s *staffRepo) GetByIdStaff(models.PrimaryKey) (models.Staff, error) {
	return models.Staff{}, nil
}

//get list  for staff

func (s *staffRepo) GetListStaff(models.GetListRequest) (models.StaffRepo, error) {
	return models.StaffRepo{}, nil
}

//update for staff

func (s *staffRepo) UpdateStaffs(models.UpdateStaff) (string, error) {
	return "", nil
}

//delete for staff

func (s *staffRepo) DeleteStaff(models.PrimaryKey) error {
	return nil
}
