package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
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
func (s *saleRepo) CreateSales(models.CreateSale) (string, error) {
	return "", nil
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
