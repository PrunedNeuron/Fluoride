package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route creates a new router, sets it up and returns it
func Route(router chi.Router) {

	// Universal routes

	// Get index page
	router.Get("/", controller.GetIndex)
	// Get current version
	router.Get("/version", controller.GetVersion)
	// Fallback if no pattern matches
	router.NotFound(controller.NotFound)

	// Icon pack specific router
	router.Mount("/{developer}", devRouter())

}
