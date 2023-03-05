package controller

import (
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokoController interface {
	GetMyToko(c *fiber.Ctx) error
	GetAllToko(c *fiber.Ctx) error
	GetTokoByID(c *fiber.Ctx) error
	UpdateTokoByID(c *fiber.Ctx) error
}

type TokoControllerImpl struct {
	Tokousecase usecase.TokoUseCase
}

func NewTokoController(tokousecase usecase.TokoUseCase) TokoController {
	return &TokoControllerImpl{
		Tokousecase: tokousecase,
	}
}

func (a *TokoControllerImpl) GetMyToko(c *fiber.Ctx) error {
	ctx := c.Context()

	claims := c.Locals("user").(jwt.MapClaims)

	userid := fmt.Sprintf("%v", claims["sub"])

	if userid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  errors.New("user id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Tokousecase.GetMyToko(ctx, userid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *TokoControllerImpl) GetAllToko(c *fiber.Ctx) error {
	ctx := c.Context()

	filter := new(dto.FilterToko)
	if err := c.QueryParser(filter); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Tokousecase.GetAllToko(ctx, dto.FilterToko{
		Nama:  filter.Nama,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *TokoControllerImpl) GetTokoByID(c *fiber.Ctx) error {
	ctx := c.Context()

	tokoid := c.Params("tokoid")

	res, err := a.Tokousecase.GetTokoByID(ctx, tokoid)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  false,
		"message": "Failed to GET data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *TokoControllerImpl) UpdateTokoByID(c *fiber.Ctx) error {
	ctx := c.Context()

	tokoid := c.Params("tokoid")
	tokoData := new(dto.TokoUpdateReq)

	if err := c.BodyParser(&tokoData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Tokousecase.UpdateTokoByID(ctx, tokoid, *tokoData)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  false,
		"message": "Failed to GET data",
		"errors":  nil,
		"data":    res,
	})
}
