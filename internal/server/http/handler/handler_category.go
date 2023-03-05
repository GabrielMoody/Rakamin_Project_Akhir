package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewCategoryRepository(containerConf.Mysqldb)
	usecase := usecase.NewCategoryUseCase(repo)
	controller := controller.NewCategoryController(usecase)

	categoryAPI := r.Group("/category", utils.MiddlewareJWTAuthorization)
	categoryAPI.Get("", controller.GetAllCategory)
	categoryAPI.Get("/:categoryid", controller.GetCategoryByID)
	categoryAPI.Post("", controller.CreateCategory)
	categoryAPI.Put("/:categoryid", controller.UpdateCategoryByID)
	categoryAPI.Delete("/:categoryid", controller.DeleteCategoryByID)
}
