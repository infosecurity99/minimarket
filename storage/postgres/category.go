package postgres

import (
	"context"
	"fmt"
	"time"

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

func (s *categoryRepo) execWithLog(query string, args ...interface{}) error {
	_, err := s.db.Exec(context.Background(), query, args...)
	if err != nil {
		fmt.Println("error during query execution:", err.Error())
	}
	return err
}

func (c *categoryRepo) CreateCategory(createCategory models.CreateCategory) (string, error) {
	uid := uuid.New()
	create_ats := time.Now()

	query := `
        INSERT INTO category VALUES ($1, $2, $3)
        `

	if err := c.execWithLog(query,
		uid,
		createCategory.Name,
		create_ats,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (c *categoryRepo) GetByIdCategory(pKey models.PrimaryKey) (models.Category, error) {
	category := models.Category{}

	query := `
        SELECT id, name, create_at
        FROM category
        WHERE id = $1
    `

	if err := c.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&category.ID,
		&category.Name,
		&category.Create_at,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.Category{}, err
	}

	return category, nil
}

func (c *categoryRepo) GetListCategory(request models.GetListRequest) (models.CategoryResponse, error) {
	var (
		categories = []models.Category{}
		count      = 0
		query      string
		page       = request.Page
		offset     = (page - 1) * request.Limit
		search     = request.Search
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
         SELECT id, name, create_at
         FROM category
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
			&category.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.CategoryResponse{}, err
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
        SET name = $1
        WHERE id = $2
    `

	if err := c.execWithLog(query,
		request.Name,
		request.ID,
	); err != nil {
		return "", err
	}

	return request.ID, nil
}

func (c *categoryRepo) DeleteCategory(request models.PrimaryKey) error {
	query := `
       DELETE FROM category
       WHERE id = $1
`
	if err := c.execWithLog(query, request.ID); err != nil {
		return err
	}
	return nil
}
