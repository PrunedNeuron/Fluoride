package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// DevRouter routes the endpoints associated with icon pack developers
func devRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/packs", controller.GetPacksByDev)
	router.Get("/packs/count", controller.GetPackCountByDev)
	router.Get("/icons/count", controller.GetIconCountByDev)

	// Packs
	router.Mount("/{pack}", packRouter())

	return router
}
