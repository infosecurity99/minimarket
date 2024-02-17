package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"connected/api/models"
	"connected/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.ICategory {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) CreateCategory(createCategory models.CreateCategory) (string, error) {
	uid := uuid.New()

	if _, err := c.db.Exec(context.Background(), `
        INSERT INTO category 
        VALUES ($1, $2, $3, $4)
    `,
		uid,
		createCategory.Name,
		createCategory.Parent_id,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (c *categoryRepo) GetByIdCategory(pKey models.PrimaryKey) (models.Category, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	category := models.Category{}

	query := `
        SELECT id, name, parent_id, created_at, updated_at
        FROM category  
        WHERE id = $1   and deleted_at = 0
    `

	if err := c.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&category.ID,
		&category.Name,
		&category.Parent_id,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		fmt.Println("error while scanning category", err.Error())
		return models.Category{}, err
	}

	if createdAt.Valid {
		category.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		category.UpdatedAt = updatedAt.String
	}

	return category, nil
}

func (c *categoryRepo) GetListCategory(request models.GetListRequest) (models.CategoryResponse, error) {
	var (
		categories           = []models.Category{}
		count                = 0
		query                string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
        SELECT COUNT(1) FROM category
`

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := c.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of category", err.Error())
		return models.CategoryResponse{}, err
	}

	query = `
         SELECT id, name, parent_id, create_at  FROM category and deleted_at = 0 
`

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.CategoryResponse{}, err
	}

	for rows.Next() {
		category := models.Category{}

		if err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.Parent_id,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.CategoryResponse{}, err
		}
		if createdAt.Valid {
			category.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			category.UpdatedAt = updatedAt.String
		}

		categories = append(categories, category)
	}

	return models.CategoryResponse{
		Categories: categories,
		Count:      count,
	}, nil
}

func (c *categoryRepo) UpdateCategory(request models.UpdateCategory) (string, error) {
	query := `
        UPDATE category
        SET name = $1,parent_id=$2, updated_at = now()
        WHERE id = $3
    `

	if _, err := c.db.Exec(context.Background(), query,

		request.Name,
		request.Parent_id,
		request.ID); err != nil {
		fmt.Println("error while updating transaction data", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (c *categoryRepo) DeleteCategory(request models.PrimaryKey) error {
	query := `update category set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := c.db.Exec(context.Background(), query, request.ID); err != nil {
		fmt.Println("error while deleting category  by id", err.Error())
		return err
	}
	return nil
}
