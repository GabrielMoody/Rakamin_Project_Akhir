package controller

import (
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityController interface {
	GetListProvincies(c *fiber.Ctx) error
	GetListCities(c *fiber.Ctx) error
	GetDetailProvinceByID(c *fiber.Ctx) error
	GetDetailCityByID(c *fiber.Ctx) error
}

type ProvinceCityControllerImpl struct {
	provinceusecase usecase.ProvinceCityUseCase
}

func NewProvinceCityController(provinceusecase usecase.ProvinceCityUseCase) ProvinceCityController {
	return &ProvinceCityControllerImpl{
		provinceusecase: provinceusecase,
	}
}

func (a *ProvinceCityControllerImpl) GetListProvincies(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.provinceusecase.GetListProvincies(ctx)

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

func (a *ProvinceCityControllerImpl) GetListCities(c *fiber.Ctx) error {
	ctx := c.Context()
	provinceid := c.Params("provinceid")

	res, err := a.provinceusecase.GetListCities(ctx, provinceid)

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

func (a *ProvinceCityControllerImpl) GetDetailProvinceByID(c *fiber.Ctx) error {
	ctx := c.Context()
	provinceid := c.Params("provinceid")

	res, err := a.provinceusecase.GetDetailProvinceByID(ctx, provinceid)

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

func (a *ProvinceCityControllerImpl) GetDetailCityByID(c *fiber.Ctx) error {
	ctx := c.Context()
	cityid := c.Params("cityid")

	res, err := a.provinceusecase.GetDetailCityByID(ctx, cityid)

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
