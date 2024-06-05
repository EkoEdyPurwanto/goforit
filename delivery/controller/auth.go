package controller

import (
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/EkoEdyPurwanto/goforit/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	UseCase usecase.AuthUseCase
	Engine  *fiber.App
}

// Route auth
func (u *AuthController) Route() {
	api := u.Engine.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/auth/register", u.registerHandler)
}

func (u *AuthController) registerHandler(ctx *fiber.Ctx) error {
	var request req.AuthRegisterRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	if err := u.UseCase.Register(request); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Registration successful"})
}

// NewUsersController Constructor
func NewUsersController(useCase usecase.AuthUseCase, engine *fiber.App) *AuthController {
	return &AuthController{
		UseCase: useCase,
		Engine:  engine,
	}
}
