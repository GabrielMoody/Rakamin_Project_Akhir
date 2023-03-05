package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(r fiber.Router, containerConf *container.Container) {
	trxRepo := repository.NewTransactionRepository(containerConf.Mysqldb)
	alamatRepo := repository.NewAlamatRepository(containerConf.Mysqldb)
	usecase := usecase.NewTransactionUseCase(trxRepo, alamatRepo)
	controller := controller.NewTransactionController(usecase)

	transactionAPI := r.Group("/trx", utils.MiddlewareJWTAuthorization)

	transactionAPI.Get("", controller.GetAllTransaction)
	transactionAPI.Get("/:transactionid", controller.GetTransactionByID)
	transactionAPI.Post("", controller.CreateTransaction)
}
