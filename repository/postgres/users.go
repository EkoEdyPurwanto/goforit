package postgres

import "database/sql"

type (
	UsersRepository interface {
		Save()
	}
	usersRepository struct {
		db *sql.DB
	}
)

func (u *usersRepository) Save() {
	//TODO implement me
	panic("implement me")
}

// NewUsersRepository Constructor
func NewUsersRepository(db *sql.DB) UsersRepository  {
	return &usersRepository{
		db: db,
	}
}
