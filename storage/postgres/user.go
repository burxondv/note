package postgres

import (
	"github.com/burxondv/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(note *repo.User) (*repo.User, error) {
	return nil, nil
}

func (ur *userRepo) Get(id int64) (*repo.User, error) {
	return nil, nil
}

func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	return nil, nil
}

func (ur *userRepo) Update(note *repo.User) (*repo.User, error) {
	return nil, nil
}

func (ur *userRepo) Delete(id int64) error {
	return nil
}
