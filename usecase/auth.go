package usecase

import (
	"time"

	"github.com/EkoEdyPurwanto/goforit/model"
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/repository/postgres"
	"github.com/EkoEdyPurwanto/goforit/utility/common"
	"github.com/EkoEdyPurwanto/goforit/utility/security"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	AuthUseCase interface {
		Register(payload req.AuthRegisterRequest) error
	}
	authUseCase struct {
		Repository postgres.AuthRepository
		Validate   *validator.Validate
		Logger     *logrus.Logger
	}
)

func (u *authUseCase) Register(payload req.AuthRegisterRequest) error {
	// struct validation
	err := u.Validate.Struct(payload)
	if err != nil {
		u.Logger.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}
	// hash password use DefaultCost
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		u.Logger.Warnf("failed hasing password: %+v", err)
		return fiber.ErrInternalServerError
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
		u.Logger.Warnf("failed save user : %+v", err)
		return fiber.ErrInternalServerError
	}
	return nil
}

// NewAuthUseCase Constructor
func NewAuthUseCase(repository postgres.AuthRepository, validate *validator.Validate, logger *logrus.Logger) AuthUseCase {
	return &authUseCase{
		Repository: repository,
		Validate:   validate,
		Logger:     logger,
	}
}
