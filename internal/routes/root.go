package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route attaches routes to the given router
func Route(router *chi.Router) {

	// Universal routes
	(*router).Get("/", controller.GetIndex())
	(*router).Get("/version", controller.GetVersion())

	// Icon pack specific router
	(*router).Mount("/{pack}", packRouter())

}
