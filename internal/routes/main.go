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
	(*router).Get("/{pack}/icons", controller.GetAllIcons)
	(*router).Get("/{pack}/icons/pending", controller.GetPendingIcons)
	(*router).Get("/{pack}/icons/done", controller.GetDoneIcons)
	(*router).Get("/{pack}/icons/{component}", controller.GetIconByComponent)
	(*router).Post("/{pack}/icons", controller.SaveIcons)
	(*router).Get("/{pack}/icons/count", controller.GetIconCount)
	(*router).Get("/{pack}/icons/pending/count", controller.GetPendingIconCount)
	(*router).Get("/{pack}/icons/done/count", controller.GetDoneIconCount)
	(*router).Put("/{pack}/icons/status", controller.UpdateStatus)
}
