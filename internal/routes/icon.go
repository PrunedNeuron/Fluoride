package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// IconRouter routes the endpoints associated with icon requests
/*
	/developers/{developer}/packs/{pack}/icons
*/
func iconRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controller.GetIconsByPackByDev)
	router.Post("/", controller.SaveIcons)
	router.Get("/{component}", controller.GetIconByComponentByPackByDev)
	router.Get("/pending", controller.GetPendingIconsByPackByDev)
	router.Get("/done", controller.GetDoneIconsByPackByDev)
	router.Put("/status", controller.UpdateIconStatus)

	return router
}
