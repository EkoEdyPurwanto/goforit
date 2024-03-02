package controller

import (
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/gofiber/fiber/v3"
)

type UsersController struct {
	usersUC usecase.UsersUseCase
	engine  *fiber.App
}

// Route users
func (u *UsersController) Route() {

}

// NewUsersController Constructor
func NewUsersController(usersUC usecase.UsersUseCase, engine *fiber.App) *UsersController {
	return &UsersController{
		usersUC: usersUC,
		engine:  engine,
	}
}
