package controller

import (
	"fmt"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	GetMyProfile(c *fiber.Ctx) error
	UpdateMyProfile(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	usersusecase usecase.UserUseCase
}

func NewUserController(userusecase usecase.UserUseCase) UserController {
	return &UserControllerImpl{
		usersusecase: userusecase,
	}
}

func (a *UserControllerImpl) GetMyProfile(c *fiber.Ctx) error {
	ctx := c.Context()

	claims := c.Locals("user").(jwt.MapClaims)

	userid := fmt.Sprintf("%v", claims["sub"])

	res, err := a.usersusecase.GetMyProfie(ctx, userid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *UserControllerImpl) UpdateMyProfile(c *fiber.Ctx) error {
	ctx := c.Context()

	claims := c.Locals("user").(jwt.MapClaims)

	userid := fmt.Sprintf("%v", claims["sub"])

	data := new(dto.ProfileReqUpdate)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.usersusecase.UpdateMyProfile(ctx, userid, *data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    res,
	})
}
