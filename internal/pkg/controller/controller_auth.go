package controller

import (
	"log"
	"time"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	LoginUser(c *fiber.Ctx) error
	RegisterNewUser(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authusecase usecase.AuthUseCase
}

func NewAuthController(authusecase usecase.AuthUseCase) AuthController {
	return &AuthControllerImpl{
		authusecase: authusecase,
	}
}

func (a *AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	ctx := c.Context()
	user := new(dto.AuthLogin)

	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to Login",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.authusecase.LoginUser(ctx, *user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to Login",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    res,
		Path:     "/",
		MaxAge:   int(time.Hour) * 1,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to Login",
		"errors":  nil,
		"data":    res,
	})
}

func (a *AuthControllerImpl) RegisterNewUser(c *fiber.Ctx) error {
	ctx := c.Context()
	user := new(dto.AuthRegister)

	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to Register new user",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.authusecase.SingUpNewUser(ctx, *user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to Register new user",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to Register new user",
		"errors":  nil,
		"data":    res,
	})
}
