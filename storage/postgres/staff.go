package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"connected/api/models"
	"connected/storage"

	"github.com/google/uuid"
)

type staffRepo struct {
	db *sql.DB
}

func NewStaffRepo(db *sql.DB) storage.IStaff {
	return &staffRepo{
		db: db,
	}
}

func (s *staffRepo) execWithLog(query string, args ...interface{}) error {
	_, err := s.db.Exec(query, args...)
	if err != nil {
		fmt.Println("error during query execution:", err.Error())
	}
	return err
}

func (s *staffRepo) CreateStaff(createStaff models.CreateStaff) (string, error) {
	uid := uuid.New()
	create_ats := time.Now()

	query := `
		INSERT INTO staff VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	if err := s.execWithLog(query,
		uid,
		createStaff.Branch_id,
		createStaff.Tarif_id,
		createStaff.Type_Stuff_Enum,
		createStaff.Name,
		createStaff.Balance,
		createStaff.Age,
		createStaff.BirthDate,
		createStaff.Login,
		createStaff.Password,
		create_ats,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (s *staffRepo) GetByIdStaff(pKey models.PrimaryKey) (models.Staff, error) {
	staff := models.Staff{}

	query := `
		SELECT id, branch_id, tarif_id, staff_type, name, balance, age, birthdate, login, password, create_at
		FROM staff
		WHERE id = $1
	`

	if err := s.db.QueryRow(query, pKey.ID).Scan(
		&staff.ID,
		&staff.Branch_id,
		&staff.Tarif_id,
		&staff.Type_Stuff_Enum,
		&staff.Name,
		&staff.Balance,
		&staff.Age,
		&staff.BirthDate,
		&staff.Login,
		&staff.Password,
		&staff.Create_at,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.Staff{}, err
	}

	return staff, nil
}

func (s *staffRepo) GetListStaff(request models.GetListRequest) (models.StaffRepo, error) {
	var (
		staffs = []models.Staff{}
		count  = 0
		query  string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery := `
		SELECT COUNT(1) FROM staff
	`

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := s.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of staff", err.Error())
		return models.StaffRepo{}, err
	}

	query = `
		SELECT id, branch_id, tarif_id, name, staff_type, balance, age, birthdate, login, password, create_at
		FROM staff
	`

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.StaffRepo{}, err
	}

	for rows.Next() {
		staff := models.Staff{}

		if err = rows.Scan(
			&staff.ID,
			&staff.Branch_id,
			&staff.Tarif_id,
			&staff.Name,
			&staff.Type_Stuff_Enum,
			&staff.Balance,
			&staff.Age,
			&staff.BirthDate,
			&staff.Login,
			&staff.Password,
			&staff.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.StaffRepo{}, err
		}

		staffs = append(staffs, staff)
	}

	return models.StaffRepo{
		Staffs: staffs,
		Count:  count,
	}, nil
}

func (s *staffRepo) UpdateStaffs(request models.UpdateStaff) (string, error) {
	query := `
		UPDATE staff
		SET branch_id = $1, tarif_id = $2, name=$3,  balance=$4 ,  age=$5 , birthdate=$6
		WHERE id = $7
	`

	if err := s.execWithLog(query,
		request.Branch_id,
		request.Tarif_id,
		request.Name,
		request.Balance,
		request.Age,
		request.BirthDate,
		request.ID,
	); err != nil {
		return "", err
	}

	return request.ID, nil
}

func (s *staffRepo) DeleteStaff(request models.PrimaryKey) error {
	query := `
		DELETE FROM staff
		WHERE id = $1
	`

	if err := s.execWithLog(query, request.ID); err != nil {
		return err
	}

	return nil
}
