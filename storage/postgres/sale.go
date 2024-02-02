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
	createAt := time.Now()

	if _, err := s.db.Exec(context.Background(), `
        INSERT INTO sale (id, branch_id, shopassistant_id, cashier_id, payment_type, price, status_type, clientname, create_at)
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
		createAt,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getby id

func (s *saleRepo) GetByIdSales(pKey models.PrimaryKey) (models.Sale, error) {
	sale := models.Sale{}

	query := `
           SELECT id, branch_id, shopassistant_id, cashier_id,payment_type, price,status_type, clientname,create_at
           FROM sale
           WHERE id = $1
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
		&sale.Create_at,
	); err != nil {
		fmt.Println("error while scanning sale", err.Error())
	}

	return sale, nil
}

// get list
func (s *saleRepo) GetListSales(request models.GetListRequestSale) (models.SaleRepos, error) {
	var (
		sales  = []models.Sale{}
		count  = 0
		query  string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery := `
        SELECT COUNT(1) FROM sale
    `

	if search != "" {
		countQuery += fmt.Sprintf(` WHERE clientname ILIKE '%%%s%%'`, search)
	}

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of transaction", err.Error())
		return models.SaleRepos{}, err
	}

	query = `
	SELECT id, branch_id, shopassistant_id, cashier_id,payment_type, price,status_type, clientname,create_at
	FROM sale
    `

	if search != "" {
		query += fmt.Sprintf(` WHERE clientname ILIKE '%%%s%%'`, search)
	}

	if request.FromPrice > 0 || request.ToPrice > 0 {

		if search == "" {
			query += " WHERE "
		} else {
			query += " AND "
		}

		// Add condition for FromAmount and ToAmount
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
			&sale.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.SaleRepos{}, nil
		}

		sales = append(sales, sale)
	}

	return models.SaleRepos{
		Sales: sales,
		Count:        count,
	}, nil
}

// update for sale

func (s *saleRepo) UpdateSales(updates models.UpdateSale) (string, error) {
	query := `
	update sale
	   set branch_id = $1, shopassistant_id = $2, cashier_id = $3, price = $4,clientname=$5
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
	query := `
          delete from sale
             where id = $1
    `
	if _, err := s.db.Exec(context.Background(), query, pKey.ID); err != nil {
		fmt.Println("error while deleting slaes  by id", err.Error())
		return err
	}

	return nil
}
