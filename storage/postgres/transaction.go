package postgres

import (
	"context"
	"fmt"
	"time"

	"connected/api/models"
	"connected/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) storage.ITransaction {
	return &transactionRepo{
		db: db,
	}
}

// create tansaction
func (t *transactionRepo) CreateTransaction(createTransaction models.CreateTransaction) (string, error) {
	uid := uuid.New()
	createAt := time.Now()
	if _, err := t.db.Exec(context.Background(), `insert into
          transaction values ($1, $2, $3, $4, $5, $6, $7, $8)
          `,
		uid,
		createTransaction.Sale_id,
		createTransaction.Staff_id,
		createTransaction.Transaction_type,
		createTransaction.Sourcetype,
		createTransaction.Amount,
		createTransaction.Description,
		createAt,
	); err != nil {
		fmt.Println("error while inserting transaction data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyid transaction

func (t *transactionRepo) GetByIdTransaction(pKey models.PrimaryKey) (models.Transaction, error) {
	transaction := models.Transaction{}

	query := `
       select id, sale_id, staff_id, transaction_type, sourcetype, amount, description, create_at from transaction where id = $1
`
	if err := t.db.QueryRow(context.Background(), query, pKey.ID).Scan(
		&transaction.ID,
		&transaction.Sale_id,
		&transaction.Staff_id,
		&transaction.Transaction_type,
		&transaction.Sourcetype,
		&transaction.Amount,
		&transaction.Description,
		&transaction.Create_at,
	); err != nil {
		fmt.Println("error while scanning transaction", err.Error())
		return models.Transaction{}, err
	}

	return transaction, nil
}

// getlisttransaction
func (t *transactionRepo) GetListTransaction(request models.GetListRequestTransaction) (models.TransactionRepo, error) {
	var (
		transactions = []models.Transaction{}
		count        = 0
		query        string
		page         = request.Page
		offset       = (page - 1) * request.Limit
		search       = request.Search
	)

	countQuery := `
        SELECT COUNT(1) FROM transaction
    `

	if search != "" {
		countQuery += fmt.Sprintf(` WHERE product_id = '%s'`, search)
	}

	if err := t.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of transaction", err.Error())
		return models.TransactionRepo{}, err
	}

	query = `
        SELECT id, sale_id, staff_id, transaction_type, sourcetype,
        amount, description, create_at FROM transaction
    `

	if search != "" {
		query += fmt.Sprintf(` WHERE product_id = '%s'`, search)
	}

	if request.FromAmount > 0 || request.ToAmount > 0 {

		if search == "" {
			query += " WHERE "
		} else {
			query += " AND "
		}

		// Add condition for FromAmount and ToAmount
		if request.FromAmount > 0 && request.ToAmount > 0 {
			query += fmt.Sprintf(`amount BETWEEN %f AND %f`, request.FromAmount, request.ToAmount)
		} else if request.FromAmount > 0 {
			query += fmt.Sprintf(`amount >= %f`, request.FromAmount)
		} else if request.ToAmount > 0 {
			query += fmt.Sprintf(`amount <= %f`, request.ToAmount)
		}
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := t.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.TransactionRepo{}, err
	}

	for rows.Next() {
		transaction := models.Transaction{}

		if err = rows.Scan(
			&transaction.ID,
			&transaction.Sale_id,
			&transaction.Staff_id,
			&transaction.Transaction_type,
			&transaction.Sourcetype,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TransactionRepo{}, nil
		}

		transactions = append(transactions, transaction)
	}

	return models.TransactionRepo{
		Transactions: transactions,
		Count:        count,
	}, nil
}

// update  transaction
func (t *transactionRepo) UpdateTransaction(updateTransaction models.UpdateTransaction) (string, error) {
	query := `
          update transaction
             set sale_id = $1, staff_id = $2, amount = $3, description = $4
                where id = $5`

	if _, err := t.db.Exec(context.Background(), query,
		updateTransaction.Sale_id,
		updateTransaction.Staff_id,
		updateTransaction.Amount,
		updateTransaction.Description,
		updateTransaction.ID); err != nil {
		fmt.Println("error while updating transaction data", err.Error())
		return "", err
	}

	return updateTransaction.ID, nil
}

// delete transaction
func (t *transactionRepo) DeleteTransaction(pKey models.PrimaryKey) error {
	query := `
          delete from transaction
             where id = $1
    `
	if _, err := t.db.Exec(context.Background(), query, pKey.ID); err != nil {
		fmt.Println("error while deleting transaction by id", err.Error())
		return err
	}

	return nil
}
