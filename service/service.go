package service

import (
	"connected/storage"
)

type IServiseManger interface {
	Branch() branchServise
	Category() categoryServise
	Product() productServise
	Sale() saleServise
	Staff_Tarif() stafftarifServise
	Staff() staffServise
	Storages() storagesServise
	Basket() basketServise
	Transaction() transactionServise
	TransactionStorage() transactionstorageServise
}

type Services struct {
	branchServise             branchServise
	categoryServise           categoryServise
	productServise            productServise
	saleServise               saleServise
	stafftarifServise         stafftarifServise
	storagesServise           storagesServise
	staffServise              staffServise
	basketServise             basketServise
	transactionServise        transactionServise
	transactionstorageServise transactionstorageServise
}

func New(storage storage.IStorage) Services {
	services := Services{}

	services.branchServise = NewBranchServise(storage)
	services.categoryServise = NewCategoryServise(storage)
	services.productServise = NewProductServise(storage)
	services.saleServise = NewSaleServise(storage)
	services.stafftarifServise = NewStafTarifServise(storage)
	services.storagesServise = NewStoragesServise(storage)
	services.staffServise = NewstaffServise(storage)
	services.basketServise = NewbasketServise(storage)
	services.transactionServise = NewtransactionServise(storage)
	services.transactionstorageServise = NewtransactionstorageServise(storage)
	return services
}

func (s Services) Branch() branchServise {
	return s.branchServise
}
func (s Services) Category() categoryServise {
	return s.categoryServise
}

func (s Services) Product() productServise {
	return s.productServise
}

func (s Services) Sale() saleServise {
	return s.saleServise
}
func (s Services) Staff_Tarif() stafftarifServise {
	return s.stafftarifServise
}

func (s Services) Staff() staffServise {
	return staffServise(s.stafftarifServise)
}

func (s Services) Storages() storagesServise {
	return s.storagesServise
}

func (s Services) Basket() basketServise {
	return s.basketServise
}

func (s Services) Transaction() transactionServise {
	return s.transactionServise
}

func (s Services) TransactionStorage() transactionstorageServise {
	return s.transactionstorageServise
}
