package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
)

type repoBranch struct {
	db *sql.DB
}

// Create implements storage.IBranchStorage.
func (repoBranch) Create(models.CreateBranch) (string, error) {
	panic("unimplemented")
}

// Update implements storage.IBranchStorage.
func (repoBranch) Update(models.UpdateBranch) (string, error) {
	panic("unimplemented")
}

func NewBranchRepo(db *sql.DB) storage.IBranchStorage {
	return &repoBranch{
		db: db,
	}
}

//create  branch

func (b *repoBranch) CreateBranch(models.CreateBranch) (string, error) {
	return "", nil
}

//getbyid branch

func (b *repoBranch) GetByID(models.PrimaryKey) (models.Branch, error) {
	return models.Branch{}, nil
}

//getlistbranch

func (b *repoBranch) GetList(models.GetListRequest) (models.BranchResponse, error) {
	return models.BranchResponse{}, nil
}

//updatebranch

func (b *repoBranch) UpdateBranch(models.UpdateBranch) (string, error) {
	return "", nil
}

//delete branch

func (b *repoBranch) Delete(models.PrimaryKey) error {
	return nil
}
