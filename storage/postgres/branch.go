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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.IBranchStorage {
	return &branchRepo{
		db: db,
	}
}

// create branch
func (b *branchRepo) Create(createBranch models.CreateBranch) (string, error) {
	uid := uuid.New()

	if _, err := b.db.Exec(context.Background(), `
		insert into   branch values ($1, $2, $3)
		`,
		uid,
		createBranch.Name,
		createBranch.Address,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

// getbyid branch
func (b *branchRepo) GetByID(pKey models.PrimaryKey) (models.Branch, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	branch := models.Branch{}

	query := `
		SELECT id, name, address,  created_at, updated_at  FROM branch WHERE id = $1 and deleted_at = 0
	`
	if err := b.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&createdAt,
		&updatedAt,
	); err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}

// getlistbranch
func (b *branchRepo) GetList(request models.GetListRequest) (models.BranchResponse, error) {
	var (
		branches             []models.Branch
		count                int
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
		SELECT COUNT(1) FROM branch   where deleted_at = 0 `

	query := `
		SELECT id, name, address, created_at, updated_at from branch  and deleted_at = 0`

	// Common logic for adding search condition to queries
	addSearchCondition := func(baseQuery string) string {
		if request.Search != "" {
			return fmt.Sprintf("%s AND (name ILIKE '%%%s%%')", baseQuery, request.Search)
		}
		return baseQuery
	}

	countQuery = addSearchCondition(countQuery)

	if err := b.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		return models.BranchResponse{}, err
	}

	query = addSearchCondition(query) + ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(context.Background(), query, request.Limit, (request.Page-1)*request.Limit)
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
			&createdAt,
			&updatedAt,
		); err != nil {
			return models.BranchResponse{}, err
		}
		if createdAt.Valid {
			branch.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.String
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
		SET name = $1, address = $2,updated_at = now()
		WHERE id = $3
	`

	if _, err := b.db.Exec(context.Background(), query, request.Name, request.Address, request.ID); err != nil {
		return "", err
	}

	return request.ID, nil
}

// delete branch
func (b *branchRepo) Delete(request models.PrimaryKey) error {

	query := `update branch set deleted_at = extract(epoch from current_timestamp) where id = $1`
	if _, err := b.db.Exec(context.Background(), query, request.ID); err != nil {
		return err
	}

	return nil
}
