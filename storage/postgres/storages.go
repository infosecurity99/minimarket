package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"fmt"
	"time"

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
	createAt := time.Now()
	if _, err := s.db.Exec(context.Background(), `
		INSERT INTO storage VALUES ($1, $2, $3, $4 ,$5)
		`,
		uid,
		request.Product_id,
		request.Branch_id,
		request.Count,
		createAt,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

// getbyid storage
func (s *storageRepo) GetByIdStorages(pKey models.PrimaryKey) (models.Storage, error) {
	stoges := models.Storage{}

	query := `
		SELECT id,product_id,branch_id, count,create_at FROM storage WHERE id = $1
	`
	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&stoges.ID,
		&stoges.Product_id,
		&stoges.Branch_id,
		&stoges.Count,
		&stoges.Create_at,
	); err != nil {
		return models.Storage{}, err
	}

	return stoges, nil
}

// getlist storage
func (s *storageRepo) GetListStorages(reuqest models.GetListRequest) (models.StorageRepos, error) {
	var (
		storages []models.Storage
		count    int
	)

	countQuery := `
		SELECT COUNT(1) FROM storage`

	query := `
		SELECT id,product_id,branch_id, count, create_at FROM   storage `

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
			&storage.Create_at,
		); err != nil {
			return models.StorageRepos{}, err
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
	SET product_id = $1, branch_id = $2 ,count=$3
	WHERE id = $4
`

	if _, err := s.db.Exec(context.Background(), query, request.Product_id, request.Branch_id, request.Count, request.ID); err != nil {
		return "", err
	}

	return request.ID, nil
}

// delete storage
func (s *storageRepo) DeleteStorages(request models.PrimaryKey) error {
	query := `
		DELETE FROM storage 
		WHERE id = $1
	`
	if _, err := s.db.Exec(context.Background(), query, request.ID); err != nil {
		return err
	}

	return nil
}
