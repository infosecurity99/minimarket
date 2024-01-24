package storage

import "connected/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
}

//  for  user interface
type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetByID(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}
