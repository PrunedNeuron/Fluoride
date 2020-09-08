package server

import (
	"icon-requests/api"
	"icon-requests/server"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

// IconServer is the API web server
type IconServer struct {
	logger    *zap.SugaredLogger
	router    chi.Router
	iconStore api.IconStore
}

// SetupIconServer will set up the API listener and attach routes
func SetupIconServer(router chi.Router, iconStore api.IconStore) error {
	server := &IconServer{
		logger:    zap.S().With("package", "server"),
		router:    router,
		iconStore: iconStore,
	}

	server.router.Get("/icons", server.GetIcons())

	return nil
}

// GetIcons responds with a list of all the icons
func (serv *IconServer) GetIcons() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := serv.iconStore.GetIcons(request.Context())
		if err != nil {
			render.Render(writer, request, server.ErrInvalidRequest(err))
			return
		}

		render.JSON(writer, request, list)
	}
}
