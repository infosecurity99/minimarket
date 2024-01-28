package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type saleRepo struct {
	db *sql.DB
}

func NewSaleRepo(db *sql.DB) storage.ISaleStorage {
	return &saleRepo{
		db: db,
	}
}

//create sale   for   sale
// create sale for sale
func (s *saleRepo) CreateSales(createSale models.CreateSale) (string, error) {
	uid := uuid.New()
	createAt := time.Now()
	if _, err := s.db.Exec(`
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

func (s *saleRepo) GetByIdSales(models.PrimaryKey) (models.Sale, error) {
	return models.Sale{}, nil
}

//get list

func (s *saleRepo) GetListSales(models.GetListRequest) (models.SaleRepo, error) {
	return models.SaleRepo{}, nil
}

// update for sale

func (s *saleRepo) UpdateSales(models.UpdateSale) (string, error) {
	return "", nil
}

//delete for sale

func (s *saleRepo) DeleteSales(models.PrimaryKey) error {
	return nil
}
