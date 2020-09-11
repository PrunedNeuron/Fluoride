package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route attaches routes to the given router
func Route(router *chi.Router) {
	(*router).Get("/", controller.GetIndex())
	(*router).Get("/version", controller.GetVersion())

	// Icons
	(*router).Get("/icons", controller.GetIcons)
	(*router).Get("/icons/{component}", controller.GetIconByComponent)
	(*router).Post("/icons", controller.SaveIcons)
	(*router).Get("/icons/count", controller.GetCount)
	(*router).Put("/icons/status", controller.UpdateStatus)
}
