package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
)

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) storage.ITransaction {
	return &transactionRepo{
		db: db,
	}
}

//create tansaction
func (t *transactionRepo) CreateTransaction(models.CreateTransaction) (string, error) {
	return "", nil
}

//getbyid transaction

func (t *transactionRepo) GetByIdTransaction(models.PrimaryKey) (models.Transaction, error) {
	return models.Transaction{}, nil
}

//getlisttransaction
func (t *transactionRepo) GetListTransaction(models.GetListRequest) (models.TransactionRepo, error) {
	return models.TransactionRepo{}, nil
}

//update  transaction
func (t *transactionRepo) UpdateTransaction(models.UpdateTransaction) (string, error) {
	return "", nil
}

//delete transaction

func (t *transactionRepo) DeleteTransaction(models.PrimaryKey) error {
	return nil
}
