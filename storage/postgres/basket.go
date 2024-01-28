package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"connected/api/models"
	"connected/storage"
	"github.com/google/uuid"
)

type basketRepo struct {
	db *sql.DB
}

func NewBasketRepo(db *sql.DB) storage.IBasket {
	return &basketRepo{
		db: db,
	}
}

func (b *basketRepo) execWithLog(query string, args ...interface{}) error {
	_, err := b.db.Exec(query, args...)
	if err != nil {
		fmt.Println("error while executing query", err.Error())
	}
	return nil
}

func (b *basketRepo) CreateBasket(createBasket models.CreateBasket) (string, error) {
	uid := uuid.New()
	create_ats := time.Now()

	query := `
        INSERT INTO basket VALUES ($1, $2, $3, $4, $5, $6)
        `

	if err := b.execWithLog(query,
		uid,
		createBasket.Sale_id,
		createBasket.Product_id,
		createBasket.Quantity,
		createBasket.Price,
		create_ats,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (b *basketRepo) GetByIdBasket(id models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}

	query := `
           SELECT id, sale_id, product_id, quantity, price, create_at
           FROM basket
           WHERE id = $1
           `

	if err := b.db.QueryRow(query, id.ID).Scan(
		&basket.ID,
		&basket.Sale_id,
		&basket.Product_id,
		&basket.Quantity,
		&basket.Price,
		&basket.Create_at,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
	}

	return basket, nil
}

func (b *basketRepo) GetListBasket(request models.GetListRequest) (models.BasketResponse, error) {
	var (
		baskets = []models.Basket{}
		count   = 0
		query   string
		page    = request.Page
		offset  = (page - 1) * request.Limit
		search  = request.Search
	)

	countQuery := `
                SELECT COUNT(1) FROM basket
                `

	if search != "" {
		countQuery += fmt.Sprintf(` AND (name ILIKE '%%%s%%')`, search)
	}

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of basket", err.Error())
		return models.BasketResponse{}, err
	}

	query = `
             SELECT id, sale_id, product_id, quantity, price, create_at
             FROM basket
             `

	if search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%%%s%%') `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		basket := models.Basket{}

		if err = rows.Scan(
			&basket.ID,
			&basket.Sale_id,
			&basket.Product_id,
			&basket.Quantity,
			&basket.Price,
			&basket.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.BasketResponse{}, nil
		}

		baskets = append(baskets, basket)
	}

	return models.BasketResponse{
		Baskets: baskets,
		Count:   count,
	}, nil
}

func (b *basketRepo) UpdateBasket(updateBasket models.UpdateBasket) (string, error) {
	query := `
         UPDATE basket
         SET quantity = $1
         WHERE id = $2
         `

	if err := b.execWithLog(query,
		updateBasket.Quantity,
		updateBasket.ID,
	); err != nil {
		return "", err
	}

	return updateBasket.ID, nil
}

func (b *basketRepo) DeleteBasket(id models.PrimaryKey) error {
	query := `
         DELETE FROM basket
         WHERE id = $1
         `

	if err := b.execWithLog(query, id.ID); err != nil {
		return err
	}

	return nil
}
