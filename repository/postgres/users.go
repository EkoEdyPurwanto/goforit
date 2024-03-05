package postgres

import (
	"database/sql"
	"github.com/EkoEdyPurwanto/goforit/model"
)

type (
	UsersRepository interface {
		Save(user model.User) error
	}
	usersRepository struct {
		db *sql.DB
	}
)

func (u *usersRepository) Save(user model.User) error {
	//TODO implement me
	panic("implement me")
}

// NewUsersRepository Constructor
func NewUsersRepository(db *sql.DB) UsersRepository  {
	return &usersRepository{
		db: db,
	}
}
