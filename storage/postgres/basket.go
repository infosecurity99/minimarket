package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type basketRepo struct {
	db *pgxpool.Pool
}

func NewBasketRepo(db *pgxpool.Pool) storage.IBasket {
	return &basketRepo{
		db: db,
	}
}

func (b *basketRepo) CreateBasket(createBasket models.CreateBasket) (string, error) {
	uid := uuid.New()
	createAt := time.Now()

	/*price, err := b.StartSell(createBasket.Product_id)
	if err != nil {
		fmt.Println("getbyidproductprice error")
		return "", err
	}

	totalPrice := price * float64(createBasket.Quantity)

	fmt.Println("Total Price:", totalPrice)*/

	if _, err := b.db.Exec(context.Background(), `
        INSERT INTO basket (id, sale_id, product_id, quantity, price, create_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `,
		uid,
		createBasket.Sale_id,
		createBasket.Product_id,
		createBasket.Quantity,
		createBasket.Price,
		createAt,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (b *basketRepo) GetByIdBasket(pKey models.PrimaryKey) (models.Basket, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}

	basket := models.Basket{}

	query := `
           SELECT  id, sale_id ,product_id,quantity,price, create_at
           FROM basket
           WHERE id = $1
           `

	if err := b.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&basket.ID,
		&basket.Sale_id,
		&basket.Product_id,
		&basket.Quantity,
		&basket.Price,
		&createdAt,
		&updatedAt,
	); err != nil {
		fmt.Println("error while scanning sale", err.Error())
	}
	if createdAt.Valid {
		basket.CreatedAt = createdAt.Time
	}

	if updatedAt.Valid {
		basket.UpdatedAt = updatedAt.String
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
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `SELECT COUNT(1) FROM basket  where  deleted_at = 0`

	if search != "" {
		countQuery += fmt.Sprintf(` WHERE sale_id = '%s'`, search)
	}

	if err := b.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of basket", err.Error())
		return models.BasketResponse{}, err
	}

	query = `SELECT id, sale_id ,product_id,quantity,price,  created_at, updated_at FROM basket  where deleted_at = 0`

	if search != "" {
		query += fmt.Sprintf(` WHERE sale_id = '%s'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		//	fmt.Println("error while querying rows", err.Error())
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
			&createdAt,
			&updatedAt,
		); err != nil {
			//	fmt.Println("error while scanning row", err.Error())
			return models.BasketResponse{}, nil
		}
		if createdAt.Valid {
			basket.CreatedAt = createdAt.Time
		}

		if updatedAt.Valid {
			basket.UpdatedAt = updatedAt.String
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
	update basket
	   set sale_id=$1 ,product_id=$2,quantity=$3,price=$4, updated_at = now()
		  where id = $5`

	if _, err := b.db.Exec(context.Background(), query,

		updateBasket.Sale_id,
		updateBasket.Product_id,
		updateBasket.Quantity,
		updateBasket.Price,
		updateBasket.ID); err != nil {
		fmt.Println("error while updating basket data", err.Error())
		return "", err
	}

	return updateBasket.ID, nil
}

func (b *basketRepo) DeleteBasket(pKey models.PrimaryKey) error {

	query := `update basket set deleted_at = extract(epoch from current_timestamp) where id = $1`


	if _, err := b.db.Exec(context.Background(), query, pKey.ID); err != nil {
		fmt.Println("error while deleting basket  by id", err.Error())
		return err
	}

	return nil
}

/*
func (p *basketRepo) StartSell(pKey string) (float64, error) {
	var price float64

	query := `
           SELECT price
           FROM product
           WHERE id = $1
`

	if err := p.db.QueryRow(context.Background(), query, pKey).Scan(
		&price,
	); err != nil {
		fmt.Println("error while scanning product  price", err.Error())
		return price, nil
	}

	return price, nil
}
*/
