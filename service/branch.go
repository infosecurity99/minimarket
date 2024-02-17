package service

import (
	"connected/api/models"
	"connected/storage"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

type branchServise struct {
	storage storage.IStorage
}

func NewBranchServise(storage storage.IStorage) branchServise {
	return branchServise{
		storage: storage,
	}
}

func (b branchServise) Create(ctx context.Context, createBranch models.CreateBranch) (models.Branch, error) {
	pKey, err := b.storage.Branch().Create(createBranch)
	if err != nil {
		fmt.Println("while error   servises layer")
	}

	branch, err := b.storage.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		fmt.Println("servises layer  branch get by id", err.Error())
	}

	return branch, nil
}

func (b branchServise) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.Branch, error) {
	branch, err := b.storage.Branch().GetByID(models.PrimaryKey{
		ID: pKey.ID,
	})
	if err != nil {
		fmt.Println("while  erorring branch servises", err.Error())
	}
	return branch, nil
}

func (b branchServise) GetList(request models.GetListRequest) (models.BranchResponse, error) {
	branchResponse, err := b.storage.Branch().GetList(request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("ERROR in service layer while getting branchs list", err.Error())
			return models.BranchResponse{}, err
		}
	}

	return branchResponse, err
}

func (b branchServise) Update(updatebranch models.UpdateBranch) (models.Branch, error) {
	pKey, err := b.storage.Branch().Update(updatebranch)
	if err != nil {
		fmt.Println("ERROR in service layer while updating updateUser", err.Error())
		return models.Branch{}, err
	}

	branch, err := b.storage.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		fmt.Println("ERROR in service layer while getting user after update", err.Error())
		return models.Branch{}, err
	}

	return branch, nil
}

func (b branchServise) Delete(pKey models.PrimaryKey) error {
	err := b.storage.Branch().Delete(pKey)
	return err
}
