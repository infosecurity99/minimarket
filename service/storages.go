package service

import (
	"connected/api/models"
	"connected/storage"
)

type storagesServise struct {
	storage storage.IStorage
}

func NewStoragesServise(storage storage.IStorage) storagesServise {
	return storagesServise{
		storage: storage,
	}
}

func (c storagesServise) Create(models.CreateStorage) (models.Storage, error) {
	return models.Storage{}, nil
}

func (c storagesServise) GetByID(pkey models.PrimaryKey) (models.Storage, error) {
	return models.Storage{}, nil
}

func (c storagesServise) GetList(models.GetListRequest) (models.StorageRepos, error) {
	return models.StorageRepos{}, nil
}

func (c storagesServise) Update(models.UpdateStorage) (models.Storage, error) {
	return models.Storage{}, nil
}

func (c storagesServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
