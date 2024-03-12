package usecase

import (
	"fmt"
	"github.com/EkoEdyPurwanto/goforit/model"
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/repository/postgres"
	"github.com/EkoEdyPurwanto/goforit/util/common"
	"github.com/EkoEdyPurwanto/goforit/util/security"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	UsersUseCase interface {
		Register(payload req.RegisterUsersRequest) error
	}
	userUseCase struct {
		Repository postgres.UsersRepository
		Validate   *validator.Validate
	}
)

func (u *userUseCase) Register(payload req.RegisterUsersRequest) error {
	// struct validation
	err := u.Validate.Struct(payload)
	if err != nil {
		logrus.WithError(err).Error()
		return fmt.Errorf("failed save user")
	}
	// hash password use DefaultCost
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	//fill value of user
	user := model.User{
		Id:        common.GenerateID(),
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
	// save user
	err = u.Repository.Save(user)
	if err != nil {
		logrus.WithError(err).Error()
		return fmt.Errorf("failed save user")
	}
	return nil
}

// NewUsersUseCase Constructor
func NewUsersUseCase(repository postgres.UsersRepository, validate *validator.Validate) UsersUseCase {
	return &userUseCase{
		Repository: repository,
		Validate:   validate,
	}
}
