package service

import (
	"connected/api/models"
	"connected/storage"
)

type stafftarifServise struct {
	storage storage.IStorage
}

func NewStafTarifServise(storage storage.IStorage) stafftarifServise {
	return stafftarifServise{
		storage: storage,
	}
}

func (c stafftarifServise) Create(models.CreateStaff_Tarif) (models.Staff_Tarif, error) {
	return models.Staff_Tarif{}, nil
}

func (c stafftarifServise) GetByID(pkey models.PrimaryKey) (models.Staff_Tarif, error) {
	return models.Staff_Tarif{}, nil
}

func (c stafftarifServise) GetList(models.GetListRequest) (models.Staff_Tarif_Repo, error) {
	return models.Staff_Tarif_Repo{}, nil
}

func (c stafftarifServise) Update(models.UpdateStaff_Tarif) (models.Staff_Tarif, error) {
	return models.Staff_Tarif{}, nil
}

func (c stafftarifServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
