package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type storageRepo struct {
	db *sql.DB
}

func NewStorRepo(db *sql.DB) storage.IStorag {
	return &storageRepo{
		db: db,
	}
}

//Create storage
func (s *storageRepo) CreateStorages(request models.CreateStorage) (string, error) {
	uid := uuid.New()
	createAt := time.Now()
	if _, err := s.db.Exec(`
		INSERT INTO storge VALUES ($1, $2, $3, $4)
		`,
		uid,
		request.Product_id,
		request.Branch_id,
		createAt,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

//getbyid storage
func (s *storageRepo) GetByIdStorages(pKey models.PrimaryKey) (models.Storage, error) {
	stoges := models.Storage{}

	query := `
		SELECT id,product_id,branch_id, create_at FROM storage WHERE id = $1
	`
	if err := s.db.QueryRow(query, pKey.ID).Scan(
		&stoges.ID,
		&stoges.Product_id,
		&stoges.Branch_id,
		&stoges.Create_at,
	); err != nil {
		return models.Storage{}, err
	}

	return stoges, nil
}

//getlist storage
func (s *storageRepo) GetListStorages(reuqest models.GetListRequest) (models.StorageRepos, error) {
	var (
		storages []models.Storage
		count    int
	)

	countQuery := `
		SELECT COUNT(1) FROM storage`

	query := `
		SELECT id,product_id,branch_id, create_at FROM   storage `


	addSearchCondition := func(baseQuery string) string {
		if reuqest.Search != "" {
			return fmt.Sprintf("%s AND (product_id ILIKE '%%%s%%')", baseQuery, reuqest.Search)
		}
		return baseQuery
	}

	countQuery = addSearchCondition(countQuery)

	if err := s.db.QueryRow(countQuery).Scan(&count); err != nil {
		return models.StorageRepos{}, err
	}

	query = addSearchCondition(query) + ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(query, reuqest.Limit, (reuqest.Page-1)*reuqest.Limit)
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

//update  storage
func (s *storageRepo) UpdateStorages(request models.UpdateStorage) (string, error) {
	query := `
	UPDATE storage
	SET product_id = $1, branch_id = $2
	WHERE id = $3
`

if _, err := s.db.Exec(query, request.Product_id, request.Branch_id, request.ID); err != nil {
	return "", err
}

return request.ID, nil
}

//delete storage
func (s *storageRepo) DeleteStorages(request models.PrimaryKey) error {
	query := `
		DELETE FROM storage 
		WHERE id = $1
	`
	if _, err := s.db.Exec(query, request.ID); err != nil {
		return err
	}

	return nil
}
