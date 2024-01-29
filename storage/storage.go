package storage

import "connected/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransaction
	Staff() IStaff
	Staff_Tarif() IStaff_Tarif
	Category() ICategory
	Product() IProduct
	Basket() IBasket
	Storag() IStorag
	TransactionStorage() ITransactionStorage
}

// for  branch interface
type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetByID(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}

// for sale interface
type ISaleStorage interface {
	CreateSales(models.CreateSale) (string, error)
	GetByIdSales(models.PrimaryKey) (models.Sale, error)
	GetListSales(models.GetListRequest) (models.SaleRepos, error)
	UpdateSales(models.UpdateSale) (string, error)
	DeleteSales(models.PrimaryKey) error
}

// for transaction interface
type ITransaction interface {
	CreateTransaction(models.CreateTransaction) (string, error)
	GetByIdTransaction(models.PrimaryKey) (models.Transaction, error)
	GetListTransaction(models.GetListRequest) (models.TransactionRepo, error)
	UpdateTransaction(models.UpdateTransaction) (string, error)
	DeleteTransaction(models.PrimaryKey) error
}

//for staff interface

type IStaff interface {
	CreateStaff(models.CreateStaff) (string, error)
	GetByIdStaff(models.PrimaryKey) (models.Staff, error)
	GetListStaff(models.GetListRequest) (models.StaffRepo, error)
	UpdateStaffs(models.UpdateStaff) (string, error)
	DeleteStaff(models.PrimaryKey) error
}

// staff_tarif
type IStaff_Tarif interface {
	CreateStaff_Tarifs(models.CreateStaff_Tarif) (string, error)
	GetByIdStaff_Tarifs(models.PrimaryKey) (models.Staff_Tarif, error)
	GetListStaff_Tarifs(models.GetListRequest) (models.Staff_Tarif_Repo, error)
	UpdateStaff_Tarifs(models.UpdateStaff_Tarif) (string, error)
	DeleteStaff_Tarifs(models.PrimaryKey) error
}

// category interface
type ICategory interface {
	CreateCategory(models.CreateCategory) (string, error)
	GetByIdCategory(models.PrimaryKey) (models.Category, error)
	GetListCategory(models.GetListRequest) (models.CategoryResponse, error)
	UpdateCategory(models.UpdateCategory) (string, error)
	DeleteCategory(models.PrimaryKey) error
}

// product interface
type IProduct interface {
	CreateProduct(models.CreateProduct) (string, error)
	GetByIdProduct(models.PrimaryKey) (models.Product, error)
	GetListProduct(models.GetListRequest) (models.ProductResponse, error)
	UpdateProduct(models.UpdateProduct) (string, error)
	DeleteProduct(models.PrimaryKey) error
}

// basket interface
type IBasket interface {
	CreateBasket(models.CreateBasket) (string, error)
	GetByIdBasket(models.PrimaryKey) (models.Basket, error)
	GetListBasket(models.GetListRequest) (models.BasketResponse, error)
	UpdateBasket(models.UpdateBasket) (string, error)
	DeleteBasket(models.PrimaryKey) error
}

// storage intarface
type IStorag interface {
	CreateStorages(models.CreateStorage) (string, error)
	GetByIdStorages(models.PrimaryKey) (models.Storage, error)
	GetListStorages(models.GetListRequest) (models.StorageRepos, error)
	UpdateStorages(models.UpdateStorage) (string, error)
	DeleteStorages(models.PrimaryKey) error
}

//transaction  intarface

type ITransactionStorage interface {
	CreateTransactionStorage(models.CreateTransactionStorage) (string, error)
	GetByIdTranasactionStorage(models.PrimaryKey) (models.TransactionStorage, error)
	GetListTransactionStorage(models.GetListRequest) (models.TransactionStorageResponse, error)
	UpdateTransactionStorage(models.UpdateTransactionStorage) (string, error)
	DeleteTransactionStorage(models.PrimaryKey) error
}
