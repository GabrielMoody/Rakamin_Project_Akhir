package controller

import (
	"errors"
	"log"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	GetAllProduk(c *fiber.Ctx) error
	GetProdukByID(c *fiber.Ctx) error
	CreateProduk(c *fiber.Ctx) error
	UpdateProdukByID(c *fiber.Ctx) error
	DeleteProdukByID(c *fiber.Ctx) error
}

type ProdukControllerImpl struct {
	Produkusecase usecase.ProdukUseCase
}

func NewProdukController(produkusecase usecase.ProdukUseCase) ProdukController {
	return &ProdukControllerImpl{
		Produkusecase: produkusecase,
	}
}

func (a *ProdukControllerImpl) GetAllProduk(c *fiber.Ctx) error {
	ctx := c.Context()

	filter := new(dto.ProdukFilter)
	if err := c.QueryParser(filter); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Produkusecase.GetAllProduk(ctx, dto.ProdukFilter{
		Nama_produk: filter.Nama_produk,
		Limit:       filter.Limit,
		Page:        filter.Page,
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

func (a *ProdukControllerImpl) GetProdukByID(c *fiber.Ctx) error {
	ctx := c.Context()
	produkid := c.Params("produkid")

	if produkid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  errors.New("produk id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Produkusecase.GetProdukByID(ctx, produkid)

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

func (a *ProdukControllerImpl) CreateProduk(c *fiber.Ctx) error {
	ctx := c.Context()

	produk := new(dto.ProdukReqCreate)
	if err := c.BodyParser(&produk); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Faield to POST data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Produkusecase.CreateProduk(ctx, *produk)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Faield to POST data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *ProdukControllerImpl) UpdateProdukByID(c *fiber.Ctx) error {
	ctx := c.Context()

	produkid := c.Params("produkid")
	produk := new(dto.ProdukReqUpdate)

	if produkid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  errors.New("produk id is empty"),
			"data":    nil,
		})
	}

	if err := c.BodyParser(&produk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Produkusecase.UpdateProdukByID(ctx, produkid, *produk)

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

func (a *ProdukControllerImpl) DeleteProdukByID(c *fiber.Ctx) error {
	ctx := c.Context()

	produkid := c.Params("produkid")

	if produkid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE data",
			"errors":  errors.New("produk id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Produkusecase.DeleteProdukByID(ctx, produkid)

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
