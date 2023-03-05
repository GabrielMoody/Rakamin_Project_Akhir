package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewRepositoryToko(containerConf.Mysqldb)
	usecase := usecase.NewTokoUseCase(repo)
	controller := controller.NewTokoController(usecase)

	tokoAPI := r.Group("/toko", utils.MiddlewareJWTAuthorization)

	tokoAPI.Get("", controller.GetAllToko)
	tokoAPI.Get("/my", controller.GetMyToko)
	tokoAPI.Get("/:tokoid", controller.GetTokoByID)
	tokoAPI.Put("/:tokoid", controller.UpdateTokoByID)
}
