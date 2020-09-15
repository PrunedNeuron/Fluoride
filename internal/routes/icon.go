package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// IconRouter routes the endpoints associated with icon requests
func iconRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controller.GetIconsByPackByDev)
	router.Post("/", controller.SaveIcons)

	router.Get("/{component}", controller.GetIconByComponentByPackByDev)

	router.Get("/count", controller.GetIconCountByDev)

	router.Get("/pending", controller.GetPendingIconsByPackByDev)
	router.Get("/pending/count", controller.GetPendingIconCountByDev)

	router.Get("/done", controller.GetDoneIconsByPackByDev)
	router.Get("/done/count", controller.GetDoneIconCountByDev)

	router.Put("/status", controller.UpdateIconStatus)

	return router
}
