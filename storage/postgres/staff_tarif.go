package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"connected/api/models"
	"connected/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type staftarifRepo struct {
	db *pgxpool.Pool
}

func NewStaff_Tarif(db *pgxpool.Pool) storage.IStaff_Tarif {
	return &staftarifRepo{
		db: db,
	}
}

// create stafftarif
func (s *staftarifRepo) CreateStaff_Tarifs(createStaff_tarif models.CreateStaff_Tarif) (string, error) {
	uid := uuid.New()

	if _, err := s.db.Exec(context.Background(),
		`INSERT INTO staff_tarif (id, name, tarif_type, amount_for_cash, amount_for_card)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		uid,
		createStaff_tarif.Name,
		createStaff_tarif.Tarif_Type_Enum,
		createStaff_tarif.Amount_For_Cashe,
		createStaff_tarif.Amount_For_Card,
	); err != nil {
		log.Printf("Error while inserting data: %v", err)
		return "", err
	}

	return uid.String(), nil
}

// getbyid  staftarif
func (s *staftarifRepo) GetByIdStaff_Tarifs(pKey models.PrimaryKey) (models.Staff_Tarif, error) {
	stafftarif1 := models.Staff_Tarif{}
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	query := `
    SELECT id, name, tarif_type, amount_for_cash, amount_for_card, created_at, updated_at  FROM staff_tarif WHERE id = $1
`

	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&stafftarif1.ID,
		&stafftarif1.Name,
		&stafftarif1.Tarif_Type_Enum,
		&stafftarif1.Amount_For_Cashe,
		&stafftarif1.Amount_For_Card,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		log.Printf("Error while scanning user: %v", err)
		return models.Staff_Tarif{}, err
	}
	if createdAt.Valid {
		stafftarif1.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		stafftarif1.UpdatedAt = updatedAt.String
	}

	return stafftarif1, nil
}

// get list
func (s *staftarifRepo) GetListStaff_Tarifs(request models.GetListRequest) (models.Staff_Tarif_Repo, error) {
	var (
		stafftarifs          = []models.Staff_Tarif{}
		count                = 0
		countQuery, query    string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery = `SELECT count(1) from staff_tarif and deleted_at = 0  `

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		log.Printf("Error while scanning count of branch: %v", err)
		return models.Staff_Tarif_Repo{}, err
	}

	query = `
		SELECT id, name, tarif_type, amount_for_cash, amount_for_card, created_at, updated_at
		FROM  staff_tarif  where   deleted_at = 0
		`

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		log.Printf("Error while querying rows: %v", err)
		return models.Staff_Tarif_Repo{}, err
	}

	for rows.Next() {
		staftarif := models.Staff_Tarif{}

		if err = rows.Scan(
			&staftarif.ID,
			&staftarif.Name,
			&staftarif.Tarif_Type_Enum,
			&staftarif.Amount_For_Cashe,
			&staftarif.Amount_For_Card,
			&createdAt,
			&updatedAt,
		); err != nil {
			log.Printf("Error while scanning row: %v", err)
			return models.Staff_Tarif_Repo{}, err
		}

		stafftarifs = append(stafftarifs, staftarif)
	}

	return models.Staff_Tarif_Repo{
		Staff_Tarif_Repos: stafftarifs,
		Count:             count,
	}, nil
}

// update list
func (s *staftarifRepo) UpdateStaff_Tarifs(request models.UpdateStaff_Tarif) (string, error) {
	query := `
		UPDATE staff_tarif
		SET name = $1,  amount_for_cash=$2, amount_for_card=$3,updated_at = now()
		WHERE id = $4
	`

	if _, err := s.db.Exec(context.Background(), query, request.Name, request.Amount_For_Cashe, request.Amount_For_Card, request.ID); err != nil {
		log.Printf("Error while updating branch data: %v", err)
		return "", err
	}

	return request.ID, nil
}

// delete staf tarif
func (s *staftarifRepo) DeleteStaff_Tarifs(request models.PrimaryKey) error {
	query := `update staff_tarif set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := s.db.Exec(context.Background(), query, request.ID); err != nil {
		log.Printf("Error while deleting branch by id: %v", err)
		return err
	}

	return nil
}
