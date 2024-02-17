package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"fmt"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.IProduct {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(createProduct models.CreateProduct) (string, error) {
	uid := uuid.New()

	query := `
        insert into product values ($1, $2, $3, $4, $5 )
        `

	if _, err := p.db.Exec(context.Background(), query,
		uid,
		createProduct.Name,
		createProduct.Price,
		createProduct.Barcode,
		createProduct.Category_id,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (p *productRepo) GetByIdProduct(pKey models.PrimaryKey) (models.Product, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	product := models.Product{}

	query := `
           SELECT id, name, price, barcode, category_id, created_at, updated_at 
           FROM product
           WHERE id = $1  and deleted_at = 0
`

	if err := p.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.Category_id,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		fmt.Println("error while scanning product", err.Error())
		return models.Product{}, nil
	}

	if createdAt.Valid {
		product.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.String
	}

	return product, nil
}

func (p *productRepo) GetListProduct(request models.GetListRequest) (models.ProductResponse, error) {
	var (
		products             = []models.Product{}
		count                = 0
		query                string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
                SELECT COUNT(1) FROM product and deleted_at = 0
`

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := p.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of product", err.Error())
		return models.ProductResponse{}, err
	}

	query = `
             SELECT id, name, price, barcode, category_id, created_at, updated_at
             FROM product and deleted_at = 0
	`

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := p.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.ProductResponse{}, err
	}

	for rows.Next() {
		product := models.Product{}

		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.Category_id,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.ProductResponse{}, nil
		}
		if createdAt.Valid {
			product.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			product.UpdatedAt = updatedAt.String
		}

		products = append(products, product)
	}

	return models.ProductResponse{
		Products: products,
		Count:    count,
	}, nil
}

func (p *productRepo) UpdateProduct(updateProduct models.UpdateProduct) (string, error) {
	query := `
         UPDATE product
         SET name = $1, price = $2, category_id = $3, updated_at = now()
         WHERE id = $4
`
	if _, err := p.db.Exec(context.Background(), query,
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.Category_id,
		updateProduct.ID,
	); err != nil {
		return "", err
	}

	return updateProduct.ID, nil
}

func (p *productRepo) DeleteProduct(id models.PrimaryKey) error {

	query := `update product set deleted_at = extract(epoch from current_timestamp) where id = $1`
	if _, err := p.db.Exec(context.Background(), query, id.ID); err != nil {
		return err
	}
	return nil
}
