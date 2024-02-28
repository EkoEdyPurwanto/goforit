package manager

import "goforit/repository/postgres"

type (
	RepositoryManager interface {
		UsersRepository() postgres.UsersRepository
	}
	repositoryManager struct {
		infraM InfraManager
	}
)

// UsersRepository implement from RepositoryManager
func (r *repositoryManager) UsersRepository() postgres.UsersRepository {
	return postgres.NewUsersRepository(r.infraM.Conn())
}

// NewRepositoryManager Constructor
func NewRepositoryManager(infraM InfraManager) RepositoryManager {
	return &repositoryManager{
		infraM: infraM,
	}
}
