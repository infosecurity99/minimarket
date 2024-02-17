package service

import (
	"connected/api/models"
	"connected/storage"
)

type saleServise struct {
	storage storage.IStorage
}

func NewSaleServise(storage storage.IStorage) saleServise {
	return saleServise{
		storage: storage,
	}
}

func (s saleServise) Create(models.CreateSale) (models.Sale, error) {
	return models.Sale{}, nil
}

func (s saleServise) GetByID(pKey models.PrimaryKey) (models.Sale, error) {
	return models.Sale{}, nil
}

func (s saleServise) GetList(request models.GetListRequest) (models.SaleRepos, error) {
	return models.SaleRepos{}, nil
}
func (s saleServise) Update(request models.UpdateSale) (models.Sale, error) {
	return models.Sale{}, nil
}

func (s saleServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
