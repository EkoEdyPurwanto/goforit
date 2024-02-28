package controller

import (
	"github.com/gofiber/fiber/v3"
	"goforit/usecase"
)

type UsersController struct {
	usersUC usecase.UsersUseCase
	engine *fiber.App
}

func (u *UsersController) Route() {

}

// NewUsersController Constructor
func NewUsersController(usersUC usecase.UsersUseCase, engine *fiber.App) *UsersController {
	return &UsersController{
		usersUC: usersUC,
		engine: engine,
	}
}
