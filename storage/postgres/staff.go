package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"connected/api/models"
	"connected/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) storage.IStaff {
	return &staffRepo{
		db: db,
	}
}

func (s *staffRepo) CreateStaff(createStaff models.CreateStaff) (string, error) {
	uid := uuid.New()

	birthDate, err := time.Parse("2006-01-02", createStaff.BirthDate)
	if err != nil {
		log.Println("Error parsing birth date:", err)
		return "", err
	}
	age := int(time.Since(birthDate).Hours() / 24 / 365)

	query := `
		INSERT INTO staff VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	if _, err := s.db.Exec(context.Background(), query,
		uid,
		createStaff.Branch_id,
		createStaff.Tarif_id,
		createStaff.Type_stuff,
		createStaff.Name,
		createStaff.Balance,
		age,
		birthDate,
		createStaff.Login,
		createStaff.Password,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (s *staffRepo) GetByIdStaff(pKey models.PrimaryKey) (models.Staff, error) {
	staff := models.Staff{}
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}

	query := `
		SELECT id, branch_id, tarif_id, type_stuff, name, balance, age, birthdate, login, password, created_at, updated_at 
		FROM staff
		WHERE id = $1  and deleted_at = 0
	`

	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&staff.ID,
		&staff.Branch_id,
		&staff.Tarif_id,
		&staff.Type_stuff,
		&staff.Name,
		&staff.Balance,
		&staff.Age,
		&staff.BirthDate,
		&staff.Login,
		&staff.Password,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.Staff{}, err
	}

	return staff, nil
}

func (s *staffRepo) GetListStaff(request models.GetListRequest) (models.StaffRepo, error) {
	var (
		staffs               = []models.Staff{}
		count                = 0
		query                string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
		SELECT COUNT(1) FROM staff and deleted_at = 0
	`

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of staff", err.Error())
		return models.StaffRepo{}, err
	}

	query = `
		SELECT id, branch_id, tarif_id, type_stuff,name, balance, age, birthdate, login, password, created_at, updated_at
		FROM staff   where   deleted_at = 0
	`

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(context.Background(), query, request.Limit, offset)
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
			&staff.Type_stuff,
			&staff.Name,
			&staff.Balance,
			&staff.Age,
			&staff.BirthDate,
			&staff.Login,
			&staff.Password,
			&createdAt,
			&updatedAt,
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
		SET branch_id = $1, tarif_id = $2, type_stuff=$3,
		  name=$4 ,  balance=$5 , login=$6,updated_at = now()
		WHERE id = $7
	`

	if _, err := s.db.Exec(context.Background(), query,
		request.BranchID,
		request.TariffID,
		request.StaffType,
		request.Name,
		request.Balance,
		request.Login,
		request.ID,
	); err != nil {
		return "", err
	}

	return request.ID, nil
}

func (s *staffRepo) DeleteStaff(request models.PrimaryKey) error {

	query := `update staff set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := s.db.Exec(context.Background(), query, request.ID); err != nil {
		return err
	}

	return nil
}

func (s *staffRepo) GetPassword(id string) (string, error) {
	password := ""

	query := `
		select password from staff 
		                where  id = $1`

	if err := s.db.QueryRow(context.Background(), query, id).Scan(&password); err != nil {
		fmt.Println("Error while scanning password from staff", err.Error())
		return "", err
	}

	return password, nil
}

func (s *staffRepo) UpdatePassword(request models.UpdateStaffPassword) error {
	query := `
		update staff 
				set password = $1
					where id = $2 `

	if _, err := s.db.Exec(context.Background(), query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for ff", err.Error())
		return err
	}

	return nil
}
