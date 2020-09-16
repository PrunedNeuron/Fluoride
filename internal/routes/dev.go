package routes

import (
	"fluoride/internal/controller"

	"github.com/go-chi/chi"
)

// DevRouter routes the endpoints associated with icon pack developers
/*
	/developers/{developer}
*/
func devRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controller.GetDevByUsername)

	// Icon requests routes
	router.Get("/icons", controller.GetIconsByDev)
	router.Get("/icons/count", controller.GetIconCountByDev)
	router.Get("/icons/count/pending", controller.GetPendingIconCountByDev)
	router.Get("/icons/count/done", controller.GetDoneIconCountByDev)

	// Packs
	router.Get("/packs", controller.GetPacksByDev)
	router.Post("/packs", controller.CreatePack)
	router.Get("/packs/count", controller.GetPackCountByDev)

	router.Mount("/packs/{pack}", packRouter())

	return router
}
