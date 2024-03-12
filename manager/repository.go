package manager

import "github.com/EkoEdyPurwanto/goforit/repository/postgres"

type (
	RepositoryManager interface {
		UsersRepository() postgres.UsersRepository
	}
	repositoryManager struct {
		InfraM InfraManager
	}
)

// UsersRepository implement from RepositoryManager
func (r *repositoryManager) UsersRepository() postgres.UsersRepository {
	return postgres.NewUsersRepository(r.InfraM.Conn())
}

// NewRepositoryManager Constructor
func NewRepositoryManager(infraM InfraManager) RepositoryManager {
	return &repositoryManager{
		InfraM: infraM,
	}
}
