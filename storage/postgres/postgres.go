package postgres

import (
	"context"
	"fmt"
	"strings"

	"connected/config"
	"connected/storage"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB)
	poolConfig, err := pgxpool.ParseConfig(url)
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

	//migrations
	m, err := migrate.New("file:///home/zarif/Desktop/minimarket/migrations/postgres", url)
	if err != nil {
		fmt.Println("error creating new migration instance:", err)
		return nil, err
	}

	if err = m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			if err != nil {
				panic(err)
			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					panic(err)
				}
			}

			panic(err)
		}
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
