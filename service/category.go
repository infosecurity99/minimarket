package service

import (
	"connected/api/models"
	"connected/storage"
)

type categoryServise struct {
	storage storage.IStorage
}

func NewCategoryServise(storage storage.IStorage) categoryServise {
	return categoryServise{
		storage: storage,
	}
}

func (c categoryServise) Create(models.CreateCategory) (models.Category, error) {
	return models.Category{}, nil
}

func (c categoryServise) GetByID(pkey models.PrimaryKey) (models.Category, error) {
	return models.Category{}, nil
}

func (c categoryServise) GetList(models.GetListRequest) (models.CategoryResponse, error) {
	return models.CategoryResponse{}, nil
}

func (c categoryServise) Update(models.UpdateCategory) (models.Category, error) {
	return models.Category{}, nil
}

func (c categoryServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
