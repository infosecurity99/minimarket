package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
)

type branchRepo struct {
	db *sql.DB
}



func NewBranchRepo(db *sql.DB) storage.IBranchStorage {
	return &branchRepo{
		db: db,
	}
}

//create  branch

func (b *branchRepo) Create(models.CreateBranch) (string, error) {
	return "", nil
}

//getbyid branch

func (b *branchRepo) GetByID(models.PrimaryKey) (models.Branch, error) {
	return models.Branch{}, nil
}

//getlistbranch

func (b *branchRepo) GetList(models.GetListRequest) (models.BranchResponse, error) {
	return models.BranchResponse{}, nil
}

//updatebranch

func (b *branchRepo) Update(models.UpdateBranch) (string, error) {
	return "", nil
}

//delete branch

func (b *branchRepo) Delete(models.PrimaryKey) error {
	return nil
}
