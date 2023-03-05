package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, containerConf *container.Container) {
	authRrepo := repository.NewAuthRepository(containerConf.Mysqldb)
	tokoRrepo := repository.NewRepositoryToko(containerConf.Mysqldb)
	usecase := usecase.NewAuthUseCase(authRrepo, tokoRrepo)
	controller := controller.NewAuthController(usecase)

	produkAPI := r.Group("/auth")
	produkAPI.Post("/login", controller.LoginUser)
	produkAPI.Post("/register", controller.RegisterNewUser)
}
