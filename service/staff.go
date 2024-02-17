package service

import (
	"connected/api/models"
	"connected/storage"
)

type staffServise struct {
	storage storage.IStorage
}

func NewstaffServise(storage storage.IStorage) staffServise {
	return staffServise{
		storage: storage,
	}
}

func (c staffServise) Create(models.CreateStaff) (models.Staff, error) {
	return models.Staff{}, nil
}

func (c staffServise) GetByID(pkey models.PrimaryKey) (models.Staff, error) {
	return models.Staff{}, nil
}

func (c staffServise) GetList(models.GetListRequest) (models.StaffRepo, error) {
	return models.StaffRepo{}, nil
}

func (c staffServise) Update(models.UpdateStaff) (models.Staff, error) {
	return models.Staff{}, nil
}

func (c staffServise) Delete(pKey models.PrimaryKey) error {
	return nil
}
