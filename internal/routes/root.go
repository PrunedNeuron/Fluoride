package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route creates a new router, sets it up and returns it
func Route(router chi.Router) {

	// Routes to create new model instances
	router.Mount("/create", creatorRouter())

	/* Universal routes */
	// Get index page
	router.Get("/", controller.GetIndex)
	// Get current version
	router.Get("/version", controller.GetVersion)

	// Icon pack specific router
	router.Mount("/{dev}", devRouter())

	// Fallback if no pattern matches
	router.NotFound(controller.NotFound)

}

func creatorRouter() chi.Router {
	router := chi.NewRouter()

	router.Post("/user", controller.CreateUser)
	router.Post("/pack", controller.CreatePack)
	router.Post("/plan", controller.CreatePlan)

	return router
}
