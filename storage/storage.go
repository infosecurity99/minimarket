package storage

type IStorage interface {
	Close()
	User() IUserStorage
}

//  for  user interface
type IUserStorage interface {
	/*Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (string, error)
	Delete(models.PrimaryKey) error
	GetPassword(id string) (string, error)
	UpdatePassword(password models.UpdateUserPassword) error*/
}
