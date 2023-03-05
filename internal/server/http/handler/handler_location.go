package handler

import (
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func LocationHandler(r fiber.Router) {
	usecase := usecase.NewProvinceCityUseCase()
	controller := controller.NewProvinceCityController(usecase)

	locationAPI := r.Group("/provcity")
	locationAPI.Get("listprovincies", controller.GetListProvincies)
	locationAPI.Get("listcities/:provinceid", controller.GetListCities)
	locationAPI.Get("detailprovince/:provinceid", controller.GetDetailProvinceByID)
	locationAPI.Get("detailcity/:cityid", controller.GetDetailCityByID)
}
