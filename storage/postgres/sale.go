package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) storage.ISaleStorage {
	return &saleRepo{
		db: db,
	}
}

// create sale   for   sale
// create sale for sale
func (s *saleRepo) CreateSales(createSale models.CreateSale) (string, error) {
	uid := uuid.New()

	if _, err := s.db.Exec(context.Background(), `
        INSERT INTO sale (id, branch_id, shopassistant_id, cashier_id, payment_type, price, status_type, clientname,)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `,
		uid,
		createSale.Branch_id,
		createSale.Shopassistant_id,
		createSale.Cashier_id,
		createSale.Payment_type,
		createSale.Price,
		createSale.Status_type,
		createSale.Clientname,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getby id

func (s *saleRepo) GetByIdSales(pKey models.PrimaryKey) (models.Sale, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	sale := models.Sale{}

	query := `
           SELECT id, branch_id, shopassistant_id, cashier_id,payment_type, price,status_type, clientname,created_at, updated_at 
           FROM sale
           WHERE id = $1 and deleted_at = 0
           `

	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&sale.ID,
		&sale.Branch_id,
		&sale.Shopassistant_id,
		&sale.Cashier_id,
		&sale.Payment_type,
		&sale.Price,
		&sale.Status_type,
		&sale.Clientname,
		&createdAt, //4
		&updatedAt, //5
	); err != nil {
		fmt.Println("error while scanning sale", err.Error())
	}
	if createdAt.Valid {
		sale.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		sale.UpdatedAt = updatedAt.String
	}

	return sale, nil
}

// get list
func (s *saleRepo) GetListSales(request models.GetListRequestSale) (models.SaleRepos, error) {
	var (
		sales                = []models.Sale{}
		count                = 0
		query                string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
        SELECT COUNT(1) FROM sale and deleted_at = 0
    `

	if search != "" {
		countQuery += fmt.Sprintf(` WHERE id= '%%%s%%'`, search)
	}

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of transaction", err.Error())
		return models.SaleRepos{}, err
	}

	query = `
	SELECT id, branch_id, shopassistant_id, cashier_id,payment_type, price,status_type, clientname,created_at, updated_at
	FROM sale  where   deleted_at = 0
    `

	if search != "" {
		query += fmt.Sprintf(` WHERE id= '%%%s%%'`, search)
	}

	if request.FromPrice > 0 || request.ToPrice > 0 {

		if search == "" {
			query += " WHERE "
		} else {
			query += " AND "
		}

		if request.FromPrice > 0 && request.ToPrice > 0 {
			query += fmt.Sprintf(`amount BETWEEN %f AND %f`, request.FromPrice, request.ToPrice)
		} else if request.FromPrice > 0 {
			query += fmt.Sprintf(`amount >= %f`, request.FromPrice)
		} else if request.ToPrice > 0 {
			query += fmt.Sprintf(`amount <= %f`, request.ToPrice)
		}
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.SaleRepos{}, err
	}

	for rows.Next() {
		sale := models.Sale{}

		if err = rows.Scan(
			&sale.ID,
			&sale.Branch_id,
			&sale.Shopassistant_id,
			&sale.Cashier_id,
			&sale.Payment_type,
			&sale.Price,
			&sale.Status_type,
			&sale.Clientname,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.SaleRepos{}, nil
		}

		if createdAt.Valid {
			sale.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			sale.UpdatedAt = updatedAt.String
		}

		sales = append(sales, sale)
	}

	return models.SaleRepos{
		Sales: sales,
		Count: count,
	}, nil
}

// update for sale

func (s *saleRepo) UpdateSales(updates models.UpdateSale) (string, error) {
	query := `
	update sale
	   set branch_id = $1, shopassistant_id = $2, cashier_id = $3, price = $4,clientname=$5, updated_at = now()
		  where id = $6`

	if _, err := s.db.Exec(context.Background(), query,

		updates.Branch_id,
		updates.Shopassistant_id,
		updates.Cashier_id,
		updates.Price,
		updates.Clientname,
		updates.ID); err != nil {
		fmt.Println("error while updating transaction data", err.Error())
		return "", err
	}

	return updates.ID, nil
}

//delete for sale

func (s *saleRepo) DeleteSales(pKey models.PrimaryKey) error {

	query := `update sale set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := s.db.Exec(context.Background(), query, pKey.ID); err != nil {
		fmt.Println("error while deleting slaes  by id", err.Error())
		return err
	}

	return nil
}

func (s *saleRepo) UpdatePrice(ctx context.Context, totalSum uint64, id string) (string, error) {
	query := `UPDATE sale SET price = $1 WHERE id = $2`

	rowsAffected, err := s.db.Exec(ctx, query, totalSum, id)
	if err != nil {
		fmt.Println("error while updating sale price:", err.Error())
		return "", err
	}

	if r := rowsAffected.RowsAffected(); r == 0 {
		fmt.Println("error in rows affected: no rows updated")
		return "", errors.New("no rows updated")
	}

	return id, nil
}
