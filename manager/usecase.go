package manager

import (
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/go-playground/validator/v10"
)

type (
	UseCaseManager interface {
		AuthUseCase() usecase.AuthUseCase
	}
	useCaseManager struct {
		RepositoryM RepositoryManager
		Validate    *validator.Validate
	}
)

// AuthUseCase implement from UseCaseManager
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.RepositoryM.AuthRepository(), u.Validate)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate) UseCaseManager {
	return &useCaseManager{
		RepositoryM: repositoryM,
		Validate:    validate,
	}
}
