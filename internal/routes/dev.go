package routes

import "github.com/go-chi/chi"

// DevRouter routes the endpoints associated with icon pack developers
func devRouter() chi.Router {
	router := chi.NewRouter()

	// Packs
	router.Mount("/{pack}", packRouter())

	return router
}
