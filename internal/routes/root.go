package routes

import (
	"github.com/PrunedNeuron/Fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// Route creates a new router, sets it up and returns it
func Route(router chi.Router) {

	// Routes to create new model instances
	router.Mount("/create", creatorRouter())

	// Get index page
	router.Get("/", controller.GetIndex)
	// Get current version
	router.Get("/version", controller.GetVersion)

	// Developer routes
	router.Get("/developers", controller.GetDevs)
	router.Get("/developers/count", controller.GetDevCount)
	router.Mount("/developers/{developer}", devRouter())

	// Icon pack routes
	router.Get("/packs", controller.GetPacks)

	// Icon Request routes
	router.Get("/icons", controller.GetIcons)

	// Plan routes
	router.Get("/plans", controller.GetPlans)
	router.Post("/plans", controller.CreatePlan)

	// User routes
	router.Get("/users", controller.GetUsers)
	router.Post("/users", controller.CreateUser)

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
