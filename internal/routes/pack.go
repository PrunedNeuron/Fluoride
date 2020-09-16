package routes

import (
	"github.com/go-chi/chi"
)

/*
	/developers/{developer}/packs/{pack}
*/
func packRouter() chi.Router {
	router := chi.NewRouter()

	// Icon requests
	router.Mount("/icons", iconRouter())

	return router

}
