package postgres

import (
	"context"
	"fmt"

	"connected/config"
	"connected/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	poolConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB))
	if err != nil {
		fmt.Println("error while parsing config", err.Error())
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	return Store{
		Pool: pool,
	}, nil
}

func (s Store) Close() {
	s.Pool.Close()
}
func (s Store) Branch() storage.IBranchStorage {
	newBranch := NewBranchRepo(s.Pool)

	return newBranch
}

func (s Store) Sale() storage.ISaleStorage {
	newBranch := NewSaleRepo(s.Pool)

	return newBranch
}
func (s Store) Transaction() storage.ITransaction {
	return NewTransactionRepo(s.Pool)
}

func (s Store) Staff() storage.IStaff {
	return NewStaffRepo(s.Pool)
}

func (s Store) Staff_Tarif() storage.IStaff_Tarif {
	return NewStaff_Tarif(s.Pool)
}

func (s Store) Category() storage.ICategory {

	return NewCategoryRepo(s.Pool)
}

func (s Store) Product() storage.IProduct {
	
	return NewProductRepo(s.Pool)
}

func (s Store) Basket() storage.IBasket {
	

	return NewBasketRepo(s.Pool)
}

func (s Store) Storag() storage.IStorag {

	return NewStorRepo(s.Pool)
}

func (s Store) TransactionStorage() storage.ITransactionStorage {

	return NewTransactioStoragenRepo(s.Pool)
}
