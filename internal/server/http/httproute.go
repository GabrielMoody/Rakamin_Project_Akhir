package http

import (
	route "tugas_akhir_example/internal/server/http/handler"

	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/")

	route.BooksRoute(api, containerConf)
	route.ProdukRoute(api, containerConf)
	route.AuthRoute(api, containerConf)
	route.CategoryRoute(api, containerConf)
	route.TokoRoute(api, containerConf)
	route.AlamatRoute(api, containerConf)
	route.UserRoute(api, containerConf)
	route.TransactionRoute(api, containerConf)
	route.LocationHandler(api)
}
