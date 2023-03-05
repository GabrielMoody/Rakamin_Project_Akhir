package controller

import (
	"errors"
	"log"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAllCategory(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategoryByID(c *fiber.Ctx) error
	DeleteCategoryByID(c *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	Categoryusecase usecase.CategoryUseCase
}

func NewCategoryController(Categoryusecase usecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		Categoryusecase: Categoryusecase,
	}
}

func (a *CategoryControllerImpl) GetAllCategory(c *fiber.Ctx) error {
	ctx := c.Context()

	// filter := new(dto.CategoryRes)
	// if err := c.QueryParser(filter); err != nil {
	// 	log.Println(err)
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	res, err := a.Categoryusecase.GetAllCategory(ctx)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
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

func (a *CategoryControllerImpl) GetCategoryByID(c *fiber.Ctx) error {
	ctx := c.Context()
	categoryid := c.Params("categoryid")

	if categoryid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  "category id is empty",
			"data":    nil,
		})
	}

	res, err := a.Categoryusecase.GetCategoryByID(ctx, categoryid)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
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

func (a *CategoryControllerImpl) CreateCategory(c *fiber.Ctx) error {
	ctx := c.Context()

	category := new(dto.CategoryReq)
	if err := c.BodyParser(&category); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Categoryusecase.CreateCategory(ctx, *category)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *CategoryControllerImpl) UpdateCategoryByID(c *fiber.Ctx) error {
	ctx := c.Context()

	categoryid := c.Params("categoryid")
	category := new(dto.CategoryReq)

	if categoryid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  errors.New("category id is empty"),
			"data":    nil,
		})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Categoryusecase.UpdateCategoryByID(ctx, categoryid, *category)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to PUT data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *CategoryControllerImpl) DeleteCategoryByID(c *fiber.Ctx) error {
	ctx := c.Context()

	categoryid := c.Params("categoryid")

	if categoryid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE data",
			"errors":  errors.New("category id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Categoryusecase.DeleteCategoryByID(ctx, categoryid)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to DELETE data",
		"errors":  nil,
		"data":    res,
	})

}
