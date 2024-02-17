package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storageRepo struct {
	db *pgxpool.Pool
}

func NewStorRepo(db *pgxpool.Pool) storage.IStorag {
	return &storageRepo{
		db: db,
	}
}

// Create storage
func (s *storageRepo) CreateStorages(request models.CreateStorage) (string, error) {
	uid := uuid.New()

	if _, err := s.db.Exec(context.Background(), `
		INSERT INTO storage VALUES ($1, $2, $3, $4 )
		`,
		uid,
		request.Product_id,
		request.Branch_id,
		request.Count,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

// getbyid storage
func (s *storageRepo) GetByIdStorages(pKey models.PrimaryKey) (models.Storage, error) {
	stoges := models.Storage{}
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}

	query := `
		SELECT id,product_id,branch_id, count,created_at, updated_at  FROM storage WHERE id = $1 and deleted_at = 0
	`
	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&stoges.ID,
		&stoges.Product_id,
		&stoges.Branch_id,
		&stoges.Count,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		return models.Storage{}, err
	}
	if createdAt.Valid {
		stoges.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		stoges.UpdatedAt = updatedAt.String
	}

	return stoges, nil
}

// getlist storage
func (s *storageRepo) GetListStorages(reuqest models.GetListRequest) (models.StorageRepos, error) {
	var (
		storages             []models.Storage
		count                int
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
		SELECT COUNT(1) FROM storage  and deleted_at = 0`

	query := `
		SELECT id,product_id,branch_id, count, created_at, updated_at FROM   storage   where   deleted_at = 0`

	addSearchCondition := func(baseQuery string) string {
		if reuqest.Search != "" {
			return fmt.Sprintf("%s where (product_id='%s')", baseQuery, reuqest.Search)
		}
		return baseQuery
	}

	countQuery = addSearchCondition(countQuery)

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		return models.StorageRepos{}, err
	}

	query = addSearchCondition(query) + ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(context.Background(), query, reuqest.Limit, (reuqest.Page-1)*reuqest.Limit)
	if err != nil {
		return models.StorageRepos{}, err
	}
	defer rows.Close()

	for rows.Next() {
		storage := models.Storage{}

		if err := rows.Scan(
			&storage.ID,
			&storage.Product_id,
			&storage.Branch_id,
			&storage.Count,
			&createdAt,
			&updatedAt,
		); err != nil {
			return models.StorageRepos{}, err
		}
		if createdAt.Valid {
			storage.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			storage.UpdatedAt = updatedAt.String
		}

		storages = append(storages, storage)
	}

	return models.StorageRepos{
		Storages: storages,
		Count:    count,
	}, nil
}

// update  storage
func (s *storageRepo) UpdateStorages(request models.UpdateStorage) (string, error) {
	query := `
	UPDATE storage
	SET product_id = $1, branch_id = $2 ,count=$3,updated_at = now()
	WHERE id = $4
`

	if _, err := s.db.Exec(context.Background(), query, request.Product_id, request.Branch_id, request.Count, request.ID); err != nil {
		return "", err
	}

	return request.ID, nil
}

// delete storage
func (s *storageRepo) DeleteStorages(request models.PrimaryKey) error {

	query := `update storage set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := s.db.Exec(context.Background(), query, request.ID); err != nil {
		return err
	}

	return nil
}
