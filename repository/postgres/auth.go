package postgres

import (
	"database/sql"

	"github.com/EkoEdyPurwanto/goforit/model"
)

type (
	AuthRepository interface {
		Save(user model.User) error
	}
	authRepository struct {
		DB *sql.DB
	}
)

func (u *authRepository) Save(user model.User) error {
	SQL := `INSERT INTO "user_credential"(id, username, email, password, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
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

// NewAuthRepository Constructor
func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		DB: db,
	}
}
