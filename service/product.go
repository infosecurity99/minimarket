package service

import (
	"connected/api/models"
	"connected/storage"
)

type productServise struct {
	storage storage.IStorage
}

func NewProductServise(storage storage.IStorage) productServise {
	return productServise{
		storage: storage,
	}
}

func (p productServise) Create(models.CreateProduct) (models.Product, error) {
	return models.Product{}, nil
}

func (p productServise) GetByID(pKey models.PrimaryKey) (models.Product, error) {
	return models.Product{}, nil
}

func (p productServise) GetList(request models.GetListRequest) (models.ProductResponse, error) {
	return models.ProductResponse{}, nil
}

func (p productServise) Update(request models.UpdateProduct) (models.Product, error) {
	return models.Product{}, nil
}

func (p productServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
