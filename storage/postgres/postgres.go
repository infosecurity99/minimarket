package postgres

import (
	"database/sql"
	"fmt"

	"connected/config"
	"connected/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	return Store{
		DB: db,
	}, nil
}

func (s Store) Close() {
	s.DB.Close()
}

func (s Store) Branch() storage.IBranchStorage {
	newBranch := NewBranchRepo(s.DB)

	return newBranch
}

func (s Store) Sale() storage.ISaleStorage {
	newBranch := NewSaleRepo(s.DB)

	return newBranch
}
func (s Store) Transaction() storage.ITransaction {
	newTransaction := NewTransactionRepo(s.DB)

	return newTransaction
}

func (s Store) Staff() storage.IStaff {
	newStaff := NewStaffRepo(s.DB)

	return newStaff
}

func (s Store) Staff_Tarif() storage.IStaff_Tarif {
	newStaff_Tarif := NewStaff_Tarif(s.DB)

	return newStaff_Tarif
}

func (s Store) Category() storage.ICategory {
	newCategory := NewCategoryRepo(s.DB)

	return newCategory
}

func (s Store) Product() storage.IProduct {
	newProduct := NewProductRepo(s.DB)

	return newProduct
}

func (s Store) Basket() storage.IBasket {
	newBasket := NewBasketRepo(s.DB)

	return newBasket
}

func (s Store) Storag() storage.IStorag {
	newStorages:= NewStorRepo(s.DB)

	return newStorages
}


func (s Store) TransactionStorage() storage.ITransactionStorage {
	newTransactionStorage:=NewTransactioStoragenRepo(s.DB)

	return newTransactionStorage
}