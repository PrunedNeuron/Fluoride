package routes

import (
	"github.com/go-chi/chi"
)

func packRouter() chi.Router {
	router := chi.NewRouter()

	// Icons
	router.Mount("/icons", iconRouter())

	return router

}
