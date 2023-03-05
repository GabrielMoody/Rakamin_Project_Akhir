package controller

import (
	"errors"
	"fmt"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TransactionController interface {
	GetAllTransaction(c *fiber.Ctx) error
	GetTransactionByID(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
}

type TransactionControllerImpl struct {
	transcationusecase usecase.TransactionUseCase
}

func NewTransactionController(transcationusecase usecase.TransactionUseCase) TransactionController {
	return &TransactionControllerImpl{
		transcationusecase: transcationusecase,
	}
}

func (a *TransactionControllerImpl) GetAllTransaction(c *fiber.Ctx) error {
	ctx := c.Context()
	res, err := a.transcationusecase.GetAllTransaction(ctx)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  err.Err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  false,
		"message": "Failed to GET data",
		"errors":  nil,
		"data":    res,
	})
}

func (a *TransactionControllerImpl) GetTransactionByID(c *fiber.Ctx) error {
	ctx := c.Context()

	transactionid := c.Params("transactionid")

	if transactionid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  errors.New("transaction id is empty"),
			"data":    nil,
		})
	}

	res, err := a.transcationusecase.GetTransactionByID(ctx, transactionid)

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

func (a *TransactionControllerImpl) CreateTransaction(c *fiber.Ctx) error {
	ctx := c.Context()

	transaction := new(dto.TransactionReq)
	claims := c.Locals("user").(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["sub"])

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	res, err := a.transcationusecase.CreateTransaction(ctx, userid, *transaction)

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
