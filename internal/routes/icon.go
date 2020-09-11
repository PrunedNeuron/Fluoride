package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// IconRouter routes the endpoints associated with icon requests
func iconRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controller.GetAllIcons)
	router.Post("/", controller.SaveIcons)

	router.Get("/{component}", controller.GetIconByComponent)

	router.Get("/count", controller.GetIconCount)

	router.Get("/pending", controller.GetPendingIcons)
	router.Get("/pending/count", controller.GetPendingIconCount)

	router.Get("/done", controller.GetDoneIcons)
	router.Get("/done/count", controller.GetDoneIconCount)

	router.Put("/status", controller.UpdateStatus)

	return router
}
