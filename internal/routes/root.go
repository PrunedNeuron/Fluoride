package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route creates a new router, sets it up and returns it
func Route(router chi.Router) {

	// Universal routes
	router.Get("/", controller.GetIndex())
	router.Get("/version", controller.GetVersion())

	// Icon pack specific router
	router.Mount("/{pack}", packRouter())
}
