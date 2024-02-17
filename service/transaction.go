package service

import (
	"connected/api/models"
	"connected/storage"
)

type transactionServise struct {
	storage storage.IStorage
}

func NewtransactionServise(storage storage.IStorage) transactionServise {
	return transactionServise{
		storage: storage,
	}
}

func (c transactionServise) Create(models.CreateTransaction) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (c transactionServise) GetByID(pkey models.PrimaryKey) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (c transactionServise) GetList(models.GetListRequest) (models.TransactionRepo, error) {
	return models.TransactionRepo{}, nil
}

func (c transactionServise) Update(models.UpdateTransaction) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (c transactionServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
