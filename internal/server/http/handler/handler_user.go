package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewUserRepository(containerConf.Mysqldb)
	usecase := usecase.NewUserUseCase(repo)
	controller := controller.NewUserController(usecase)

	produkAPI := r.Group("/user", utils.MiddlewareJWTAuthorization)
	produkAPI.Get("", controller.GetMyProfile)
	produkAPI.Put("", controller.UpdateMyProfile)
}
