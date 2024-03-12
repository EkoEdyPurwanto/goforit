package controller

import (
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UsersController struct {
	UsersUC usecase.UsersUseCase
	Engine  *fiber.App
}

// Route users
func (u *UsersController) Route() {
	rg := u.Engine.Group("/api/v1")

	rg.Post("/auth/register", u.registerHandler)
}

func (u *UsersController) registerHandler(ctx *fiber.Ctx) error {
	var request req.RegisterUsersRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	if err := u.UsersUC.Register(request); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Registration successful"})
}

// NewUsersController Constructor
func NewUsersController(usersUC usecase.UsersUseCase, engine *fiber.App) *UsersController {
	return &UsersController{
		UsersUC: usersUC,
		Engine:  engine,
	}
}
