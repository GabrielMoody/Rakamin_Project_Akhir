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

type AlamatController interface {
	GetMyAlamat(c *fiber.Ctx) error
	GetAlamatByID(c *fiber.Ctx) error
	CreateAlamat(c *fiber.Ctx) error
	UpdateAlamatByID(c *fiber.Ctx) error
	DeleteAlamatByID(c *fiber.Ctx) error
}

type AlamatControllerImpl struct {
	Alamatusecase usecase.AlamatUseCase
}

func NewAlamatController(Alamatusecase usecase.AlamatUseCase) AlamatController {
	return &AlamatControllerImpl{
		Alamatusecase: Alamatusecase,
	}
}

func (a *AlamatControllerImpl) GetMyAlamat(c *fiber.Ctx) error {
	ctx := c.Context()

	claims := c.Locals("user").(jwt.MapClaims)

	userid := fmt.Sprintf("%v", claims["sub"])

	res, err := a.Alamatusecase.GetMyAlamat(ctx, userid)

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

func (a *AlamatControllerImpl) GetAlamatByID(c *fiber.Ctx) error {
	ctx := c.Context()
	alamatid := c.Params("Alamatid")

	if alamatid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  errors.New("alamat id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Alamatusecase.GetAlamatByID(ctx, alamatid)

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

func (a *AlamatControllerImpl) CreateAlamat(c *fiber.Ctx) error {
	ctx := c.Context()

	Alamat := new(dto.AlamatReq)
	if err := c.BodyParser(&Alamat); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	claims := c.Locals("user").(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["sub"])

	res, err := a.Alamatusecase.CreateAlamat(ctx, userid, *Alamat)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
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

func (a *AlamatControllerImpl) UpdateAlamatByID(c *fiber.Ctx) error {
	ctx := c.Context()

	Alamatid := c.Params("Alamatid")
	Alamat := new(dto.AlamatReq)

	if Alamatid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  errors.New("alamat id is empty"),
			"data":    nil,
		})
	}

	if err := c.BodyParser(&Alamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to PUT data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.Alamatusecase.UpdateAlamatByID(ctx, Alamatid, *Alamat)

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

func (a *AlamatControllerImpl) DeleteAlamatByID(c *fiber.Ctx) error {
	ctx := c.Context()

	Alamatid := c.Params("alamatid")

	if Alamatid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE data",
			"errors":  errors.New("alamat id is empty"),
			"data":    nil,
		})
	}

	res, err := a.Alamatusecase.DeleteAlamatByID(ctx, Alamatid)

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
