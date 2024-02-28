package usecase

import "goforit/repository/postgres"

type (
	UsersUseCase interface {
		//TODO implement me
	}
	userUseCase struct {
		repository postgres.UsersRepository
	}
)

// NewUsersUseCase Constructor
func NewUsersUseCase(repository postgres.UsersRepository) UsersUseCase {
	return &userUseCase{
		repository: repository,
	}
}
