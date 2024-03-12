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
		DB *sql.DB
	}
)

func (u *usersRepository) Save(user model.User) error {
	SQL := `INSERT INTO "user"(id, username, email, password, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := u.DB.Exec(SQL,
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// NewUsersRepository Constructor
func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		DB: db,
	}
}
