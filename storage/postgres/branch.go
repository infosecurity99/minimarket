package postgres

import (
	"connected/api/models"
	"connected/storage"

	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) storage.IBranchStorage {
	return &branchRepo{
		db: db,
	}
}

// create branch
func (b *branchRepo) Create(createBranch models.CreateBranch) (string, error) {
	uid := uuid.New()
	createAt := time.Now()
	if _, err := b.db.Exec(`
		INSERT INTO branch VALUES ($1, $2, $3, $4)
		`,
		uid,
		createBranch.Name,
		createBranch.Address,
		createAt,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

// getbyid branch
func (b *branchRepo) GetByID(pKey models.PrimaryKey) (models.Branch, error) {
	branch := models.Branch{}

	query := `
		SELECT id, name, address, create_at FROM branch WHERE id = $1
	`
	if err := b.db.QueryRow(query, pKey.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.Create_at,
	); err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}

// getlistbranch
func (b *branchRepo) GetList(request models.GetListRequest) (models.BranchResponse, error) {
	var (
		branches []models.Branch
		count    int
	)

	countQuery := `
		SELECT COUNT(1) FROM branch`

	query := `
		SELECT id, name, address, create_at FROM branch`

	// Common logic for adding search condition to queries
	addSearchCondition := func(baseQuery string) string {
		if request.Search != "" {
			return fmt.Sprintf("%s AND (name ILIKE '%%%s%%')", baseQuery, request.Search)
		}
		return baseQuery
	}

	countQuery = addSearchCondition(countQuery)

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		return models.BranchResponse{}, err
	}

	query = addSearchCondition(query) + ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		return models.BranchResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		branch := models.Branch{}

		if err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.Create_at,
		); err != nil {
			return models.BranchResponse{}, err
		}

		branches = append(branches, branch)
	}

	return models.BranchResponse{
		Branches: branches,
		Count:    count,
	}, nil
}

// updatebranch
func (b *branchRepo) Update(request models.UpdateBranch) (string, error) {
	query := `
		UPDATE branch
		SET name = $1, address = $2
		WHERE id = $3
	`

	if _, err := b.db.Exec(query, request.Name, request.Address, request.ID); err != nil {
		return "", err
	}

	return request.ID, nil
}

// delete branch
func (b *branchRepo) Delete(request models.PrimaryKey) error {
	query := `
		DELETE FROM branch
		WHERE id = $1
	`
	if _, err := b.db.Exec(query, request.ID); err != nil {
		return err
	}

	return nil
}
