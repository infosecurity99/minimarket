package postgres

import (
	"connected/api/models"
	"connected/storage"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (b *branchRepo) Create(createBranch models.CreateBranch) (string, error) {
	uid := uuid.New()
	create_at := time.Now()
	if _, err := b.db.Exec(`insert into 
			branch values ($1, $2, $3, $4)
			`,
		uid,
		createBranch.Name,
		createBranch.Address,
		create_at,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyid branch

func (b *branchRepo) GetByID(pKey models.PrimaryKey) (models.Branch, error) {
	branch := models.Branch{}

	query := `
		select id, name, address, create_at from branch where id = $1 
`
	if err := b.db.QueryRow(query, pKey.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.Create_at,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.Branch{}, err
	}

	return branch, nil
}

//getlistbranch

func (b *branchRepo) GetList(request models.GetListRequest) (models.BranchResponse, error) {
	var (
		branches          = []models.Branch{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `
		SELECT count(1) from branch  `

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%' )`, search)
	}

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of branch", err.Error())
		return models.BranchResponse{}, err
	}

	query = `
		select id, name, address, create_at
			FROM branch
			    `

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%' ) `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.BranchResponse{}, err
	}

	for rows.Next() {
		branch := models.Branch{}

		if err = rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.Create_at,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.BranchResponse{}, err
		}

		branches = append(branches, branch)
	}

	return models.BranchResponse{
		Branches: branches,
		Count:    count,
	}, nil
}

//updatebranch

func (b *branchRepo) Update(request models.UpdateBranch) (string, error) {
	query := `
		update branch 
			set name = $1, address = $2, ctreate_at = $3
				where id = $4`

	if _, err := b.db.Exec(query, request.Name, request.Address, request.Create_at, request.ID); err != nil {
		fmt.Println("error while updating branch data", err.Error())
		return "", err
	}

	return request.ID, nil
}

//delete branch

func (b *branchRepo) Delete(request models.PrimaryKey) error {
	query := `
		delete from branch
			where id = $1
`
	if _, err := b.db.Exec(query, request.ID); err != nil {
		fmt.Println("error while deleting branch by id", err.Error())
		return err
	}

	return nil
}
