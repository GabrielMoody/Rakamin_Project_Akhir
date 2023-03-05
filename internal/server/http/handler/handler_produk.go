package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func ProdukRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewProdukRepository(containerConf.Mysqldb)
	usecase := usecase.NewProdukUseCase(repo)
	controller := controller.NewProdukController(usecase)

	produkAPI := r.Group("/product", utils.MiddlewareJWTAuthorization)

	produkAPI.Get("", controller.GetAllProduk)
	produkAPI.Get("/:produkid", controller.GetProdukByID)
	produkAPI.Post("", controller.CreateProduk)
	produkAPI.Put("/:produkid", controller.UpdateProdukByID)
	produkAPI.Delete("/:produkid", controller.DeleteProdukByID)
}
