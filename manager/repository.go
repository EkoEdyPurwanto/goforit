package manager

import "github.com/EkoEdyPurwanto/goforit/repository/postgres"

type (
	RepositoryManager interface {
		AuthRepository() postgres.AuthRepository
	}
	repositoryManager struct {
		InfraM InfraManager
	}
)

// AuthRepository implement from RepositoryManager
func (r *repositoryManager) AuthRepository() postgres.AuthRepository {
	return postgres.NewAuthRepository(r.InfraM.Conn())
}

// NewRepositoryManager Constructor
func NewRepositoryManager(infraM InfraManager) RepositoryManager {
	return &repositoryManager{
		InfraM: infraM,
	}
}
