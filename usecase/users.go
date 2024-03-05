package usecase

import (
	"github.com/EkoEdyPurwanto/goforit/exception"
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/repository/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	UsersUseCase interface {
		Register(payload req.RegisterUsersRequest) error
	}
	userUseCase struct {
		repository postgres.UsersRepository
		validate   *validator.Validate
	}
)

func (u *userUseCase) Register(payload req.RegisterUsersRequest) error {
	err := u.validate.Struct(payload)
	if err != nil {
		logrus.WithError(err).Error("Validation failed")
		return &exception.MyError{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Err:     err.Error(),
		}
	}

	return nil
}

// NewUsersUseCase Constructor
func NewUsersUseCase(repository postgres.UsersRepository, validate *validator.Validate) UsersUseCase {
	return &userUseCase{
		repository: repository,
		validate: validate,
	}
}
