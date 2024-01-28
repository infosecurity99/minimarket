package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"connected/api/models"
	"connected/storage"
	"github.com/google/uuid"
)

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) storage.ITransaction {
	return &transactionRepo{
		db: db,
	}
}

// create tansaction
func (t *transactionRepo) CreateTransaction(createTransaction models.CreateTransaction) (string, error) {
	uid := uuid.New()
	create_ats := time.Now()
	if _, err := t.db.Exec(`insert into
          transaction values ($1, $2, $3, $4, $5, $6, $7, $8)
          `,
		uid,
		createTransaction.Sale_id,
		createTransaction.Staff_id,
		createTransaction.Transaction_type_enum,
		createTransaction.Source_type_enum,
		createTransaction.Amount,
		createTransaction.Description,
		create_ats,
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
       select id, sale_id, staff_id, transaction_type_enum, source_type_enum, amount, description, create_at from transaction where id = $1
`
	if err := t.db.QueryRow(query, pKey.ID).Scan(
		&transaction.ID,
		&transaction.Sale_id,
		&transaction.Staff_id,
		&transaction.Transaction_type_enum,
		&transaction.Source_type_enum,
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
func (t *transactionRepo) GetListTransaction(request models.GetListRequest) (models.TransactionRepo, error) {
	var (
		transactions      = []models.Transaction{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `
       SELECT count(1) from transaction  `

	if search != "" {
		countQuery += fmt.Sprintf(` and (description ilike '%%%s%%' )`, search)
	}

	if err := t.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of transactions", err.Error())
		return models.TransactionRepo{}, err
	}

	query = `
       select id, sale_id, staff_id, transaction_type_enum, source_type_enum, amount, description, create_at
          FROM transaction
              `

	if search != "" {
		query += fmt.Sprintf(` and (description ilike '%%%s%%' ) `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := t.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.TransactionRepo{}, err
	}

	for rows.Next() {
		transaction := models.Transaction{}

		if err = rows.Scan(
			&transaction.ID,
			&transaction.Sale_id,
			&transaction.Staff_id,
			&transaction.Transaction_type_enum,
			&transaction.Source_type_enum,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TransactionRepo{}, err
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

	if _, err := t.db.Exec(query,
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
	if _, err := t.db.Exec(query, pKey.ID); err != nil {
		fmt.Println("error while deleting transaction by id", err.Error())
		return err
	}

	return nil
}
