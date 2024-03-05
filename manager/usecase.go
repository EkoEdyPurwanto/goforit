package manager

import (
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/go-playground/validator/v10"
)

type (
	UseCaseManager interface {
		UsersUseCase() usecase.UsersUseCase
	}
	useCaseManager struct {
		repositoryM RepositoryManager
		validate    *validator.Validate
	}
)

// UsersUseCase implement from UseCaseManager
func (u *useCaseManager) UsersUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repositoryM.UsersRepository(), u.validate)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate) UseCaseManager {
	return &useCaseManager{
		repositoryM: repositoryM,
		validate: validate,
	}
}
