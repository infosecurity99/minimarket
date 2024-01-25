package storage

import "connected/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransaction
	Staff() IStaff
	Staff_Tarif() IStaff_Tarif
}

//  for  branch interface
type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetByID(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}

//for sale interface
type ISaleStorage interface {
	CreateSales(models.CreateSale) (string, error)
	GetByIdSales(models.PrimaryKey) (models.Sale, error)
	GetListSales(models.GetListRequest) (models.SaleRepo, error)
	UpdateSales(models.UpdateSale) (string, error)
	DeleteSales(models.PrimaryKey) error
}

//for transaction interface
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

//staff_tarif
type IStaff_Tarif interface {
	CreateStaff_Tarifs(models.CreateStaff_Tarif) (string, error)
	GetByIdStaff_Tarifs(models.PrimaryKey) (models.Staff_Tarif, error)
	GetListStaff_Tarifs(models.GetListRequest) (models.Staff_Tarif_Repo, error)
	UpdateStaff_Tarifs(models.UpdateStaff_Tarif) (string, error)
	DeleteStaff_Tarifs(models.PrimaryKey) error
}
