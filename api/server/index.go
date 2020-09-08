package server

import (
	"icon-requests/api"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Server is the API web server
type Server struct {
	logger    *zap.SugaredLogger
	router    chi.Router
	iconStore api.IconStore
}

// SetupIndexServer will set up the API listener and attach routes
func SetupIndexServer(router chi.Router, iconStore api.IconStore) error {
	server := &IconServer{
		logger:    zap.S().With("package", "server"),
		router:    router,
		iconStore: iconStore,
	}

	server.router.Get("/icons", server.GetIcons())

	return nil
}
