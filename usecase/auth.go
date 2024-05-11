package usecase

import (
	"fmt"
	"github.com/EkoEdyPurwanto/goforit/model"
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/repository/postgres"
	"github.com/EkoEdyPurwanto/goforit/utility/common"
	"github.com/EkoEdyPurwanto/goforit/utility/security"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	AuthUseCase interface {
		Register(payload req.AuthRegisterRequest) error
	}
	authUseCase struct {
		Repository postgres.AuthRepository
		Validate   *validator.Validate
	}
)

func (u *authUseCase) Register(payload req.AuthRegisterRequest) error {
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

// NewAuthUseCase Constructor
func NewAuthUseCase(repository postgres.AuthRepository, validate *validator.Validate) AuthUseCase {
	return &authUseCase{
		Repository: repository,
		Validate:   validate,
	}
}
