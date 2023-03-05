package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AlamatRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewAlamatRepository(containerConf.Mysqldb)
	usecase := usecase.NewAlamatUseCase(repo)
	controller := controller.NewAlamatController(usecase)

	alamatAPI := r.Group("/user/alamat", utils.MiddlewareJWTAuthorization)

	alamatAPI.Get("", controller.GetMyAlamat)
	alamatAPI.Get("/:alamatid", controller.GetAlamatByID)
	alamatAPI.Post("", controller.CreateAlamat)
	alamatAPI.Put("/:alamatid", controller.UpdateAlamatByID)
	alamatAPI.Delete("/:alamatid", controller.DeleteAlamatByID)
}
