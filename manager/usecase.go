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
		RepositoryM RepositoryManager
		Validate    *validator.Validate
	}
)

// UsersUseCase implement from UseCaseManager
func (u *useCaseManager) UsersUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.RepositoryM.UsersRepository(), u.Validate)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate) UseCaseManager {
	return &useCaseManager{
		RepositoryM: repositoryM,
		Validate: validate,
	}
}
