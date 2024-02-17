package service

import (
	"connected/api/models"
	"connected/storage"
)

type basketServise struct {
	storage storage.IStorage
}

func NewbasketServise(storage storage.IStorage) basketServise {
	return basketServise{
		storage: storage,
	}
}

func (c basketServise) Create(models.CreateBasket) (models.Basket, error) {
	return models.Basket{}, nil
}

func (c basketServise) GetByID(pkey models.PrimaryKey) (models.Basket, error) {
	return models.Basket{}, nil
}

func (c basketServise) GetList(models.GetListRequest) (models.BasketResponse, error) {
	return models.BasketResponse{}, nil
}

func (c basketServise) Update(models.UpdateBasket) (models.Basket, error) {
	return models.Basket{}, nil
}

func (c basketServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
