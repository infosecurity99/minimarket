package postgres

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionstorageRepo struct {
	db *pgxpool.Pool
}

func NewTransactioStoragenRepo(db *pgxpool.Pool) storage.ITransactionStorage {
	return &transactionstorageRepo{
		db: db,
	}
}

// create  transaction storage
func (s *transactionstorageRepo) CreateTransactionStorage(request models.CreateTransactionStorage) (string, error) {
	uid := uuid.New()

	if _, err := s.db.Exec(context.Background(), `
		INSERT INTO storage_transaction(id, branch_id, staff_id, product_id, transaction_type, price, quantity) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		uid,
		request.Branch_id,
		request.Staff_id,
		request.Product_id,
		request.Transaction_type,
		request.Price,
		request.Quantity,
	); err != nil {
		return "", err
	}

	return uid.String(), nil
}

// getbyid  transaction storage
func (s *transactionstorageRepo) GetByIdTranasactionStorage(pKey models.PrimaryKey) (models.TransactionStorage, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	tranasactionstorage := models.TransactionStorage{}

	query := `
           SELECT id, branch_id, staff_id, product_id,   transaction_type, price,quantity,created_at, updated_at   FROM storage_transaction   WHERE id = $1 and deleted_at = 0 
           `

	if err := s.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&tranasactionstorage.ID,
		&tranasactionstorage.Branch_id,
		&tranasactionstorage.Staff_id,
		&tranasactionstorage.Product_id,
		&tranasactionstorage.Product_id,
		&tranasactionstorage.Transaction_type,
		&tranasactionstorage.Price,
		&tranasactionstorage.Quantity,
		&createdAt,
		&updatedAt,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
	}
	if createdAt.Valid {
		tranasactionstorage.Create_at = createdAt.Time
	}

	if updatedAt.Valid {
		tranasactionstorage.UpdatedAt = updatedAt.String
	}
	return tranasactionstorage, nil
}

// getlist  transaction storage
func (s *transactionstorageRepo) GetListTransactionStorage(request models.GetListRequest) (models.TransactionStorageResponse, error) {
	var (
		tranasactionstorages = []models.TransactionStorage{}
		count                = 0
		query                string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery := `
                SELECT COUNT(1) FROM storage_transaction  where deleted_at = 0 
                `

	if search != "" {
		countQuery += fmt.Sprintf(` WHERE product_id = '%s'`, search)
	}

	if err := s.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of basket", err.Error())
		return models.TransactionStorageResponse{}, err
	}

	query = `
	SELECT id, branch_id, staff_id, product_id, 
	transaction_type, price,quantity,create_at
	FROM storage_transaction  where deleted_at = 0
             `

	if search != "" {
		query += fmt.Sprintf(` WHERE product_id = '%s'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.TransactionStorageResponse{}, err
	}

	for rows.Next() {
		tranasactionstorage := models.TransactionStorage{}

		if err = rows.Scan(
			&tranasactionstorage.ID,
			&tranasactionstorage.Branch_id,
			&tranasactionstorage.Staff_id,
			&tranasactionstorage.Product_id,
			&tranasactionstorage.Product_id,
			&tranasactionstorage.Transaction_type,
			&tranasactionstorage.Price,
			&tranasactionstorage.Quantity,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TransactionStorageResponse{}, nil
		}
		if createdAt.Valid {
			tranasactionstorage.Create_at = createdAt.Time
		}

		if updatedAt.Valid {
			tranasactionstorage.UpdatedAt = updatedAt.String
		}

		tranasactionstorages = append(tranasactionstorages, tranasactionstorage)
	}

	return models.TransactionStorageResponse{
		TransactionStorages: tranasactionstorages,
		Count:               count,
	}, nil
}

// update  transaction storage
func (s *transactionstorageRepo) UpdateTransactionStorage(request models.UpdateTransactionStorage) (string, error) {
	query := `
		UPDATE storage_transaction
		SET branch_id = $1, staff_id = $2,product_id=$3,price=$4,quantity=$5,updated_at = now()
		WHERE id = $6
	`

	if _, err := s.db.Exec(context.Background(), query,
		request.Branch_id,
		request.Staff_id,
		request.Product_id,
		request.Price,
		request.Quantity,
		request.ID); err != nil {
		return "", err
	}

	return request.ID, nil
}

// update  transaction storage
func (s *transactionstorageRepo) DeleteTransactionStorage(request models.PrimaryKey) error {

	query := `update storage_transaction set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := s.db.Exec(context.Background(), query, request.ID); err != nil {
		return err
	}

	return nil
}
