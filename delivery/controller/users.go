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
	u.engine.Get("/ping", u.registerHandler)
}

func (u *UsersController) registerHandler(ctx fiber.Ctx) error {
	return nil
}

// NewUsersController Constructor
func NewUsersController(usersUC usecase.UsersUseCase, engine *fiber.App) *UsersController {
	return &UsersController{
		usersUC: usersUC,
		engine:  engine,
	}
}
