package manager

import (
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	UseCaseManager interface {
		AuthUseCase() usecase.AuthUseCase
	}
	useCaseManager struct {
		RepositoryM RepositoryManager
		Validate    *validator.Validate
		Logger      *logrus.Logger
	}
)

// AuthUseCase implement from UseCaseManager
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.RepositoryM.AuthRepository(), u.Validate, u.Logger)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate, logger *logrus.Logger) UseCaseManager {
	return &useCaseManager{
		RepositoryM: repositoryM,
		Validate:    validate,
		Logger:      logger,
	}
}
