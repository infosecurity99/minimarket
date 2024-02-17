package service

import (
	"connected/api/models"
	"connected/storage"
)

type transactionstorageServise struct {
	storage storage.IStorage
}

func NewtransactionstorageServise(storage storage.IStorage) transactionstorageServise {
	return transactionstorageServise{
		storage: storage,
	}
}

func (c transactionstorageServise) Create(models.CreateTransactionStorage) (models.TransactionStorage, error) {
	return models.TransactionStorage{}, nil
}

func (c transactionstorageServise) GetByID(pkey models.PrimaryKey) (models.TransactionStorage, error) {
	return models.TransactionStorage{}, nil
}

func (c transactionstorageServise) GetList(models.GetListRequest) (models.TransactionStorageResponse, error) {
	return models.TransactionStorageResponse{}, nil
}

func (c transactionstorageServise) Update(models.UpdateTransactionStorage) (models.TransactionStorage, error) {
	return models.TransactionStorage{}, nil
}

func (c transactionstorageServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
