package manager

import "goforit/usecase"

type (
	UseCaseManager interface {
		UsersUseCase() usecase.UsersUseCase
	}
	useCaseManager struct {
		repositoryM RepositoryManager
	}
)

// UsersUseCase implement from UseCaseManager
func (u *useCaseManager) UsersUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repositoryM.UsersRepository())
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repositoryM: repositoryM,
	}
}
