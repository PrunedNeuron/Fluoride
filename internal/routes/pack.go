package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func packRouter() http.Handler {
	router := chi.NewRouter()

	// Icons
	router.Mount("/icons", iconRouter())

	return router

}
